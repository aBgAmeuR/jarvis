package daemon

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/abgameur/jarvis/internal/config"
	"github.com/abgameur/jarvis/internal/protocol"
	"github.com/abgameur/jarvis/internal/pve"
	"github.com/abgameur/jarvis/internal/secrets"
)

type Server struct {
	SocketPath  string
	ConfigDir   string
	SecretsFile string
	Timeout     time.Duration
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

	switch req.Op {
	case protocol.OpValidate:
		s.validate(conn)
	case protocol.OpListServices:
		resp := s.listServices()
		if resp.Error != "" {
			log.Printf("op=%s error: %s", req.Op, resp.Error)
		} else {
			log.Printf("op=%s ok", req.Op)
		}
		if err := json.NewEncoder(conn).Encode(resp); err != nil {
			log.Printf("encode response: %v", err)
		}
	default:
		resp := protocol.Response{Error: fmt.Sprintf("unsupported operation %q", req.Op)}
		log.Printf("op=%s error: %s", req.Op, resp.Error)
		if err := json.NewEncoder(conn).Encode(resp); err != nil {
			log.Printf("encode response: %v", err)
		}
	}
}

func (s *Server) listServices() protocol.Response {
	return protocol.Response{Services: []protocol.ServiceInfo{}}
}

func joinErrors(errs []error) string {
	parts := make([]string, 0, len(errs))
	for _, e := range errs {
		parts = append(parts, e.Error())
	}
	return strings.Join(parts, "; ")
}

func (s *Server) emit(enc *json.Encoder, layer string, errs []error) bool {
	ev := protocol.ValidateEvent{Layer: layer, OK: len(errs) == 0}
	if !ev.OK {
		ev.Error = joinErrors(errs)
	}
	if err := enc.Encode(ev); err != nil {
		log.Printf("encode validate event: %v", err)
		return false
	}
	if ev.OK {
		log.Printf("validate layer=%s ok", ev.Layer)
	} else {
		log.Printf("validate layer=%s fail: %s", ev.Layer, ev.Error)
	}
	return ev.OK
}

func (s *Server) validate(conn net.Conn) {
	enc := json.NewEncoder(conn)

	secretsPath := s.SecretsFile
	if secretsPath == "" {
		secretsPath = filepath.Join(s.ConfigDir, "secrets.yaml")
	}

	loaded := loadConfigAndSecrets(s.ConfigDir, secretsPath)
	if !s.emit(enc, protocol.LayerParse, loaded.parseErrs) {
		return
	}
	if !s.emit(enc, protocol.LayerRequired, loaded.cfg.CheckRequired()) {
		return
	}
	if !s.emit(enc, protocol.LayerValues, loaded.cfg.CheckValues()) {
		return
	}

	ref := loaded.cfg.Proxmox.Credentials.Ref
	if !s.emit(enc, protocol.LayerSecrets, loaded.stub.CheckRef(ref)) {
		return
	}
	if !s.emit(enc, protocol.LayerFingerprint, loaded.stub.CheckFingerprint(ref)) {
		return
	}

	timeout := s.Timeout
	if timeout <= 0 {
		timeout = 15 * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	runner := s.PVERunner
	if runner == nil {
		mat, err := loaded.stub.Resolve(loaded.cfg.Proxmox.Credentials.Provider, ref)
		if err != nil {
			s.emit(enc, protocol.LayerSSH, []error{err})
			return
		}
		r, err := pve.NewSSHRunner(pve.SSHConfig{
			Address:            loaded.cfg.Proxmox.Address,
			User:               loaded.cfg.Proxmox.User,
			PrivateKeyPath:     mat.PrivateKeyPath,
			HostKeyFingerprint: mat.HostKeyFingerprint,
			Timeout:            timeout,
		})
		if err != nil {
			s.emit(enc, protocol.LayerSSH, []error{err})
			return
		}
		runner = r
	}

	if err := pve.CheckSSH(ctx, runner); err != nil {
		s.emit(enc, protocol.LayerSSH, []error{err})
		return
	}
	s.emit(enc, protocol.LayerSSH, nil)

	if err := pve.GuestsExist(ctx, runner, loaded.cfg.Jarvis.VMID, loaded.cfg.Docker.VMID); err != nil {
		s.emit(enc, protocol.LayerGuests, []error{err})
		return
	}
	s.emit(enc, protocol.LayerGuests, nil)
}

type loaded struct {
	cfg       *config.Config
	stub      *secrets.Stub
	parseErrs []error
}

func loadConfigAndSecrets(configDir, secretsPath string) loaded {
	out := loaded{cfg: &config.Config{}}
	cfg, err := config.Load(configDir)
	if err != nil {
		out.parseErrs = append(out.parseErrs, err)
	} else {
		out.cfg = cfg
	}
	stub, err := secrets.LoadStub(secretsPath)
	if err != nil {
		out.parseErrs = append(out.parseErrs, err)
	} else {
		out.stub = stub
	}
	return out
}
