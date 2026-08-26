package pve

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/abgameur/jarvis/internal/protocol"
)

// dockerArgv builds the remote argv for one docker command inside the CT.
func dockerArgv(vmid int, args ...string) []string {
	argv := []string{"sudo", "-n", "/usr/sbin/pct", "exec", strconv.Itoa(vmid), "--", "docker"}
	return append(argv, args...)
}

// containerInspect is the subset of docker inspect output jarvis reads.
// State.Health is absent when the image has no healthcheck.
type containerInspect struct {
	Name  string `json:"Name"`
	State struct {
		Status string `json:"Status"`
		Health *struct {
			Status string `json:"Status"`
		} `json:"Health"`
	} `json:"State"`
}

// ListContainers returns all containers inside the docker CT, sorted by name.
func ListContainers(ctx context.Context, runner Runner, vmid int) ([]protocol.Container, error) {
	idsRaw, err := runner.Run(ctx, dockerArgv(vmid, "ps", "--all", "--quiet", "--no-trunc"))
	if err != nil {
		return nil, err
	}
	ids := strings.Fields(string(idsRaw))
	out := make([]protocol.Container, 0, len(ids))
	if len(ids) == 0 {
		return out, nil
	}

	raw, err := runner.Run(ctx, dockerArgv(vmid, append([]string{"inspect"}, ids...)...))
	if err != nil {
		return nil, err
	}
	var infos []containerInspect
	if err := json.Unmarshal(raw, &infos); err != nil {
		return nil, fmt.Errorf("parse docker inspect json: %w", err)
	}

	for _, info := range infos {
		name := strings.TrimPrefix(info.Name, "/")
		if name == "" {
			return nil, fmt.Errorf("container missing name in docker inspect output")
		}
		if info.State.Status == "" {
			return nil, fmt.Errorf("container %s missing state", name)
		}
		health := protocol.HealthNA
		if info.State.Health != nil && info.State.Health.Status != "" {
			health = info.State.Health.Status
		}
		out = append(out, protocol.Container{Name: name, State: info.State.Status, Health: health})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

// PowerContainer runs docker start or docker stop on one container by name or id.
func PowerContainer(ctx context.Context, runner Runner, vmid int, action, container string) error {
	_, err := runner.Run(ctx, dockerArgv(vmid, action, container))
	return err
}
