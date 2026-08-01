package client

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/abgameur/jarvis/internal/protocol"
)

func Call(socketPath string, req protocol.Request) (protocol.Response, error) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return protocol.Response{}, fmt.Errorf("connect to daemon at %s: %w (is jarvisd running?)", socketPath, err)
	}
	defer conn.Close()

	if err := json.NewEncoder(conn).Encode(req); err != nil {
		return protocol.Response{}, fmt.Errorf("send request: %w", err)
	}

	var resp protocol.Response
	if err := json.NewDecoder(conn).Decode(&resp); err != nil {
		return protocol.Response{}, fmt.Errorf("read response: %w", err)
	}

	if resp.Error != "" {
		return protocol.Response{}, fmt.Errorf("%s", resp.Error)
	}

	return resp, nil
}
