package daemon

import (
	"context"
	"fmt"
	"strings"

	"github.com/abgameur/jarvis/internal/protocol"
	"github.com/abgameur/jarvis/internal/pve"
)

func handleDockerLs(ctx context.Context, sess *session, _ protocol.Request) (any, error) {
	if err := requireDockerCTRunning(ctx, sess); err != nil {
		return nil, err
	}
	containers, err := pve.ListContainers(ctx, sess.runner, sess.cfg.Docker.VMID)
	if err != nil {
		return nil, err
	}
	return protocol.DockerLsResult{Containers: containers}, nil
}

func handleDockerStart(ctx context.Context, sess *session, req protocol.Request) (any, error) {
	return handleDockerPower(ctx, sess, "start", req.Container)
}

func handleDockerStop(ctx context.Context, sess *session, req protocol.Request) (any, error) {
	return handleDockerPower(ctx, sess, "stop", req.Container)
}

func handleDockerPower(ctx context.Context, sess *session, action, container string) (any, error) {
	if strings.TrimSpace(container) == "" {
		return nil, fmt.Errorf("container name or id is required")
	}
	if err := requireDockerCTRunning(ctx, sess); err != nil {
		return nil, err
	}
	return nil, pve.PowerContainer(ctx, sess.runner, sess.cfg.Docker.VMID, action, container)
}

// requireDockerCTRunning fails docker ops when the docker CT is down
// (FR-06: name the vmid, never start it here).
func requireDockerCTRunning(ctx context.Context, sess *session) error {
	vmid := sess.cfg.Docker.VMID
	status, err := pve.CTStatus(ctx, sess.runner, vmid)
	if err != nil {
		return err
	}
	if status != "running" {
		return fmt.Errorf("docker CT %d (docker.vmid) is %s; start it before running docker commands", vmid, status)
	}
	return nil
}
