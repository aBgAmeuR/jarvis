package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/abgameur/jarvis/internal/protocol"
)

// Stream sends req and yields ValidateEvent lines until EOF.
func Stream(socketPath string, req protocol.Request, handle func(protocol.ValidateEvent) error) error {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return fmt.Errorf("connect to daemon at %s: %w (is jarvisd running?)", socketPath, err)
	}
	defer conn.Close()

	if err := json.NewEncoder(conn).Encode(req); err != nil {
		return fmt.Errorf("send request: %w", err)
	}

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
