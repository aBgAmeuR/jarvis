// jarvis is the CLI. Commands live in one file per area (proxmox.go,
// validate.go); each exposes a newXxxCmd constructor wired up here.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := newRootCmd().Execute(); err != nil {
		if errors.Is(err, errValidationFailed) {
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "jarvis: %v\n", err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	var socketPath string

	root := &cobra.Command{
		Use:           "jarvis",
		Short:         "Homelab ops CLI",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.PersistentFlags().StringVar(&socketPath, "socket", "/tmp/jarvis.sock", "Unix socket path")

	root.AddCommand(newProxmoxCmd(&socketPath))
	root.AddCommand(newValidateCmd(&socketPath))
	return root
}
