package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/abgameur/jarvis/internal/protocol"
)

func main() {
	socketPath := flag.String("socket", "/tmp/jarvis.sock", "Unix socket path")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 || args[0] != "service" || args[1] != "ls" {
		fmt.Fprintf(os.Stderr, "usage: jarvis [--socket path] service ls\n")
		os.Exit(2)
	}

	if err := serviceLS(*socketPath); err != nil {
		fmt.Fprintf(os.Stderr, "jarvis: %v\n", err)
		os.Exit(1)
	}

}

func serviceLS(socketPath string) error {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return fmt.Errorf("connect to daemon at %s: %w (is jarvisd running?)", socketPath, err)
	}

	defer conn.Close()

	req := protocol.Request{Op: "list_services"}
	if err := json.NewEncoder(conn).Encode(req); err != nil {
		return fmt.Errorf("send request: %w", err)
	}

	var resp protocol.Response
	if err := json.NewDecoder(conn).Decode(&resp); err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.Error != "" {
		return fmt.Errorf("%s", resp.Error)
	}

	fmt.Printf("%-20s\n", "NAME")
	for _, service := range resp.Services {
		fmt.Printf("%-20s\n", service.Name)
	}

	return nil
}
