package daemon

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/abgameur/jarvis/internal/config"
	"github.com/abgameur/jarvis/internal/protocol"
)

type Server struct {
	SocketPath string
	ConfigDir  string
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

	resp := s.dispatch(req)
	if resp.Error != "" {
		log.Printf("op=%s error: %s", req.Op, resp.Error)
	} else {
		log.Printf("op=%s ok", req.Op)
	}

	if err := json.NewEncoder(conn).Encode(resp); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func (s *Server) dispatch(req protocol.Request) protocol.Response {
	switch req.Op {
	case protocol.OpListServices:
		return s.listServices()
	default:
		return protocol.Response{Error: fmt.Sprintf("unsupported operation %q", req.Op)}
	}
}

func (s *Server) listServices() protocol.Response {
	cfg, err := config.Load(s.ConfigDir)
	if err != nil {
		return protocol.Response{Error: fmt.Sprintf("read config: %v", err)}
	}

	services := make([]protocol.ServiceInfo, 0, len(cfg.Services))
	for _, svc := range cfg.Services {
		services = append(services, protocol.ServiceInfo{Name: svc.Name})
	}

	return protocol.Response{Services: services}
}
