package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/abgameur/jarvis/internal/client"
	"github.com/abgameur/jarvis/internal/protocol"
)

func newDockerCmd(socketPath *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docker",
		Short: "Docker containers on the docker CT",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "ls",
		Short: "List containers",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			containers, err := client.DockerLs(*socketPath)
			if err != nil {
				return err
			}
			return printContainerTable(containers)
		},
	})
	cmd.AddCommand(newContainerPowerCmd(socketPath, "start", "Start a container", client.DockerStart))
	cmd.AddCommand(newContainerPowerCmd(socketPath, "stop", "Stop a container", client.DockerStop))
	return cmd
}

func newContainerPowerCmd(socketPath *string, verb, short string, run func(socketPath, container string) error) *cobra.Command {
	return &cobra.Command{
		Use:   verb + " <container>",
		Short: short,
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return run(*socketPath, args[0])
		},
	}
}

func printContainerTable(containers []protocol.Container) error {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "NAME\tSTATE\tHEALTH")
	for _, c := range containers {
		fmt.Fprintf(tw, "%s\t%s\t%s\n", c.Name, c.State, c.Health)
	}
	return tw.Flush()
}
