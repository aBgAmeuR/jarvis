package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/abgameur/jarvis/internal/protocol"
)

func open(socketPath string, req protocol.Request) (net.Conn, error) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("connect to daemon at %s: %w (is jarvisd running?)", socketPath, err)
	}
	if err := json.NewEncoder(conn).Encode(req); err != nil {
		conn.Close()
		return nil, fmt.Errorf("send request: %w", err)
	}
	return conn, nil
}

// Stream sends req and yields ValidateEvent lines until EOF.
func Stream(socketPath string, req protocol.Request, handle func(protocol.ValidateEvent) error) error {
	conn, err := open(socketPath, req)
	if err != nil {
		return err
	}
	defer conn.Close()

	dec := json.NewDecoder(conn)
	for {
		var ev protocol.ValidateEvent
		if err := dec.Decode(&ev); err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("read event: %w", err)
		}
		if err := handle(ev); err != nil {
			return err
		}
	}
}

// ProxmoxLs asks the daemon for the guest table. One JSON response.
func ProxmoxLs(socketPath string) ([]protocol.Guest, error) {
	conn, err := open(socketPath, protocol.Request{Op: protocol.OpProxmoxLs})
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var payload struct {
		Error  string           `json:"error"`
		Guests []protocol.Guest `json:"guests"`
	}
	if err := json.NewDecoder(conn).Decode(&payload); err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("read response: empty reply")
		}
		return nil, fmt.Errorf("read response: %w", err)
	}
	if payload.Error != "" {
		return nil, fmt.Errorf("%s", payload.Error)
	}
	if payload.Guests == nil {
		payload.Guests = []protocol.Guest{}
	}
	return payload.Guests, nil
}
