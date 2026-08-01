package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/abgameur/jarvis/internal/daemon"
)

func main() {
	configDir := flag.String("config-dir", "examples/config", "path to config directory")
	socketPath := flag.String("socket", "/tmp/jarvis.sock", "Unix socket path")
	flag.Parse()

	srv := &daemon.Server{
		SocketPath: *socketPath,
		ConfigDir:  *configDir,
	}
	if err := srv.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "jarvisd: %v\n", err)
		os.Exit(1)
	}
}
