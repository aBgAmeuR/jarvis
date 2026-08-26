package daemon

import (
	"context"
	"fmt"

	"github.com/abgameur/jarvis/internal/protocol"
	"github.com/abgameur/jarvis/internal/pve"
)

func handleProxmoxLs(ctx context.Context, sess *session, _ protocol.Request) (any, error) {
	guests, err := pve.ListGuests(ctx, sess.runner)
	if err != nil {
		return nil, err
	}
	return protocol.ProxmoxLsResult{Guests: guests}, nil
}

func handleProxmoxStart(ctx context.Context, sess *session, req protocol.Request) (any, error) {
	if err := requireLXC(ctx, sess, req.VMID); err != nil {
		return nil, err
	}
	return nil, pve.PowerLXC(ctx, sess.runner, "start", req.VMID)
}

func handleProxmoxStop(ctx context.Context, sess *session, req protocol.Request) (any, error) {
	// FR-05: never pct stop the container jarvisd runs in.
	if req.VMID == sess.cfg.Jarvis.VMID {
		return nil, fmt.Errorf("refusing to stop vmid %d: jarvisd runs in this container (jarvis.vmid)", req.VMID)
	}
	if err := requireLXC(ctx, sess, req.VMID); err != nil {
		return nil, err
	}
	return nil, pve.PowerLXC(ctx, sess.runner, "stop", req.VMID)
}

// requireLXC resolves vmid on the node and rejects anything that is not an
// LXC container (FR-05: name the type in the error).
func requireLXC(ctx context.Context, sess *session, vmid int) error {
	g, err := pve.LookupGuest(ctx, sess.runner, vmid)
	if err != nil {
		return err
	}
	if g.Type != protocol.GuestTypeLXC {
		return fmt.Errorf("guest %d is type %s; start/stop supports lxc only", vmid, g.Type)
	}
	return nil
}
