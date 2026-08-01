package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/abgameur/jarvis/internal/client"
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

	if err := listServices(*socketPath); err != nil {
		fmt.Fprintf(os.Stderr, "jarvis: %v\n", err)
		os.Exit(1)
	}
}

func listServices(socketPath string) error {
	resp, err := client.Call(socketPath, protocol.Request{Op: protocol.OpListServices})
	if err != nil {
		return err
	}

	fmt.Printf("%-20s\n", "NAME")
	for _, svc := range resp.Services {
		fmt.Printf("%-20s\n", svc.Name)
	}
	return nil
}
