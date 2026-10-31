package bootstrap

import (
	"context"

	"jarvis/server/internal/domain/overview"
	"jarvis/server/internal/domain/supervision"
	"jarvis/server/internal/source/proxmox"
)

type overviewInventory struct {
	proxmox *proxmox.Client
}

func (o overviewInventory) Inventory(ctx context.Context) (overview.Inventory, error) {
	inventory, err := o.proxmox.Inventory(ctx)
	if err != nil {
		return overview.Inventory{}, err
	}
	out := overview.Inventory{Guests: overview.Guests{Total: len(inventory.Guests)}}
	for _, g := range inventory.Guests {
		if g.Status == "running" {
			out.Guests.Running++
		}
	}
	for _, n := range inventory.Nodes {
		out.Nodes = append(out.Nodes, overview.Node{Name: n.Name, Online: n.Online, Usage: nodeUsage(n.Usage)})
	}
	return out, nil
}

func nodeUsage(u *supervision.Usage) *overview.Usage {
	if u == nil {
		return nil
	}
	return &overview.Usage{
		CPUUsageRatio:    u.CPUUsageRatio,
		MemoryUsedBytes:  u.MemoryUsedBytes,
		MemoryTotalBytes: u.MemoryTotalBytes,
		DiskUsedBytes:    u.DiskUsedBytes,
		DiskTotalBytes:   u.DiskTotalBytes,
	}
}
