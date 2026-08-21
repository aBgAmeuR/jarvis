package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/abgameur/jarvis/internal/daemon"
)

func main() {
	configDir := flag.String("config-dir", "examples/config", "path to config directory")
	secretsFile := flag.String("secrets-file", "", "path to secrets YAML (default: <config-dir>/secrets.yaml)")
	socketPath := flag.String("socket", "/tmp/jarvis.sock", "Unix socket path")
	timeout := flag.Duration("timeout", 15*time.Second, "SSH / pvesh timeout")
	flag.Parse()

	secretsPath := *secretsFile
	if secretsPath == "" {
		secretsPath = filepath.Join(*configDir, "secrets.yaml")
	}

	srv := &daemon.Server{
		SocketPath:  *socketPath,
		ConfigDir:   *configDir,
		SecretsFile: secretsPath,
		Timeout:     *timeout,
	}
	if err := srv.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "jarvisd: %v\n", err)
		os.Exit(1)
	}
}
