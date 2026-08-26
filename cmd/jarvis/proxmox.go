package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/abgameur/jarvis/internal/client"
	"github.com/abgameur/jarvis/internal/protocol"
)

func newProxmoxCmd(socketPath *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxmox",
		Short: "Proxmox guests",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "ls",
		Short: "List LXC and VM guests",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			guests, err := client.ProxmoxLs(*socketPath)
			if err != nil {
				return err
			}
			return printGuestTable(guests)
		},
	})
	cmd.AddCommand(newPowerCmd(socketPath, "start", "Start an LXC container", client.ProxmoxStart))
	cmd.AddCommand(newPowerCmd(socketPath, "stop", "Stop an LXC container", client.ProxmoxStop))
	return cmd
}

func newPowerCmd(socketPath *string, verb, short string, run func(socketPath string, vmid int) error) *cobra.Command {
	return &cobra.Command{
		Use:   verb + " <vmid>",
		Short: short,
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			vmid, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("vmid must be a number, got %q", args[0])
			}
			return run(*socketPath, vmid)
		},
	}
}

func printGuestTable(guests []protocol.Guest) error {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tTYPE\tNAME\tPOWER")
	for _, g := range guests {
		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\n", g.ID, g.Type, g.Name, g.Power)
	}
	return tw.Flush()
}
