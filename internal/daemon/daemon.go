// Package daemon implements jarvisd: a Unix-socket server that dispatches
// CLI requests to Proxmox over SSH.
//
// Each unary op is one entry in the ops registry. handleConnection owns the
// wire (decode request, encode reply, log); handlers own only their op logic.
// To add an op: define it in protocol, write a handler, register it here.
package daemon

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/abgameur/jarvis/internal/protocol"
	"github.com/abgameur/jarvis/internal/pve"
)

// Remote command timeouts (NFR-01).
const (
	queryTimeout  = 15 * time.Second
	mutateTimeout = 60 * time.Second
)

// handlerFunc runs one unary op and returns its success payload (nil for an
// empty reply) or an error.
type handlerFunc func(ctx context.Context, sess *session, req protocol.Request) (any, error)

type opSpec struct {
	timeout time.Duration
	handle  handlerFunc
}

var ops = map[string]opSpec{
	protocol.OpProxmoxLs:    {timeout: queryTimeout, handle: handleProxmoxLs},
	protocol.OpProxmoxStart: {timeout: mutateTimeout, handle: handleProxmoxStart},
	protocol.OpProxmoxStop:  {timeout: mutateTimeout, handle: handleProxmoxStop},
}

type Server struct {
	SocketPath  string
	ConfigDir   string
	SecretsFile string // empty = <ConfigDir>/secrets.yaml
	// PVERunner overrides SSH for tests. nil = build SSH runner.
	PVERunner pve.Runner
}

func (s *Server) Run() error {
	if err := os.Remove(s.SocketPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove existing socket %s: %w", s.SocketPath, err)
	}

	listener, err := net.Listen("unix", s.SocketPath)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", s.SocketPath, err)
	}
	defer listener.Close()

	log.Printf("listening on %s", s.SocketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("accept connection: %w", err)
		}
		s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	var req protocol.Request
	if err := json.NewDecoder(conn).Decode(&req); err != nil {
		log.Printf("decode request: %v", err)
		return
	}

	// validate streams NDJSON events and bypasses the unary registry.
	if req.Op == protocol.OpValidate {
		s.validate(conn)
		return
	}

	spec, ok := ops[req.Op]
	if !ok {
		respond(conn, req.Op, nil, fmt.Errorf("unsupported operation %q", req.Op))
		return
	}

	sess, err := s.newSession(spec.timeout)
	if err != nil {
		respond(conn, req.Op, nil, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), spec.timeout)
	defer cancel()

	payload, err := spec.handle(ctx, sess, req)
	respond(conn, req.Op, payload, err)
}

// respond writes the single JSON reply of a unary op and logs the outcome.
func respond(conn net.Conn, op string, payload any, err error) {
	enc := json.NewEncoder(conn)
	if err != nil {
		log.Printf("op=%s error: %s", op, err)
		if encErr := enc.Encode(protocol.Response{Error: err.Error()}); encErr != nil {
			log.Printf("encode response: %v", encErr)
		}
		return
	}
	log.Printf("op=%s ok", op)
	if payload == nil {
		payload = protocol.Response{}
	}
	if encErr := enc.Encode(payload); encErr != nil {
		log.Printf("encode response: %v", encErr)
	}
}

func joinErrors(errs []error) string {
	parts := make([]string, 0, len(errs))
	for _, e := range errs {
		parts = append(parts, e.Error())
	}
	return strings.Join(parts, "; ")
}
