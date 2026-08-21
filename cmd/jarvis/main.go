package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/abgameur/jarvis/internal/client"
	"github.com/abgameur/jarvis/internal/protocol"
)

var errValidationFailed = errors.New("validation failed")

const (
	ansiReset = "\033[0m"
	ansiGreen = "\033[32m"
	ansiRed   = "\033[31m"
)

func main() {
	var socketPath string

	root := &cobra.Command{
		Use:           "jarvis",
		Short:         "Homelab ops CLI",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.PersistentFlags().StringVar(&socketPath, "socket", "/tmp/jarvis.sock", "Unix socket path")

	root.AddCommand(&cobra.Command{
		Use:   "validate",
		Short: "Validate config, secrets, and Proxmox guest ids via jarvisd",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			color := useColor()
			failed := false
			err := client.Stream(socketPath, protocol.Request{Op: protocol.OpValidate}, func(ev protocol.ValidateEvent) error {
				if ev.OK {
					fmt.Println(markOK(color), ev.Layer)
					return nil
				}
				failed = true
				fmt.Println(markFail(color), ev.Layer)
				fmt.Printf("%serror: %s\n", red(color), ev.Error)
				return nil
			})
			if err != nil {
				return err
			}
			if failed {
				return errValidationFailed
			}
			return nil
		},
	})

	serviceCmd := &cobra.Command{
		Use:   "service",
		Short: "Service commands (legacy)",
	}
	serviceCmd.AddCommand(&cobra.Command{
		Use:   "ls",
		Short: "List services (legacy v0.2)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := client.Call(socketPath, protocol.Request{Op: protocol.OpListServices})
			if err != nil {
				return err
			}
			fmt.Printf("%-20s\n", "NAME")
			for _, svc := range resp.Services {
				fmt.Printf("%-20s\n", svc.Name)
			}
			return nil
		},
	})
	root.AddCommand(serviceCmd)

	if err := root.Execute(); err != nil {
		if errors.Is(err, errValidationFailed) {
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "jarvis: %v\n", err)
		os.Exit(1)
	}
}

func useColor() bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	return term.IsTerminal(int(os.Stdout.Fd()))
}

func markOK(color bool) string {
	if color {
		return ansiGreen + "✓" + ansiReset
	}
	return "✓"
}

func markFail(color bool) string {
	if color {
		return ansiRed + "✗" + ansiReset
	}
	return "✗"
}

func red(color bool) string {
	if color {
		return ansiRed
	}
	return ""
}
