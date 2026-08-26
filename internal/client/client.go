// Package client is the CLI side of the jarvisd Unix socket protocol.
package client

import (
	"bytes"
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

// call sends req, checks the reply's error field, and decodes the payload
// into out (nil when the op has no payload).
func call(socketPath string, req protocol.Request, out any) error {
	conn, err := open(socketPath, req)
	if err != nil {
		return err
	}
	defer conn.Close()

	data, err := io.ReadAll(conn)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if len(bytes.TrimSpace(data)) == 0 {
		return fmt.Errorf("read response: empty reply")
	}
	var status protocol.Response
	if err := json.Unmarshal(data, &status); err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if status.Error != "" {
		return fmt.Errorf("%s", status.Error)
	}
	if out != nil {
		if err := json.Unmarshal(data, out); err != nil {
			return fmt.Errorf("read response: %w", err)
		}
	}
	return nil
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

// ProxmoxLs asks the daemon for the guest table.
func ProxmoxLs(socketPath string) ([]protocol.Guest, error) {
	var res protocol.ProxmoxLsResult
	if err := call(socketPath, protocol.Request{Op: protocol.OpProxmoxLs}, &res); err != nil {
		return nil, err
	}
	if res.Guests == nil {
		res.Guests = []protocol.Guest{}
	}
	return res.Guests, nil
}

// ProxmoxStart starts an LXC container by vmid.
func ProxmoxStart(socketPath string, vmid int) error {
	return call(socketPath, protocol.Request{Op: protocol.OpProxmoxStart, VMID: vmid}, nil)
}

// ProxmoxStop stops an LXC container by vmid.
func ProxmoxStop(socketPath string, vmid int) error {
	return call(socketPath, protocol.Request{Op: protocol.OpProxmoxStop, VMID: vmid}, nil)
}
