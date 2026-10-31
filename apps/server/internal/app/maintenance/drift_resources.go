package maintenance

import (
	"context"
	"fmt"

	"jarvis/server/internal/domain/maintenance"
)

func (d *Drift) resourceMismatches(ctx context.Context, inProxmox, inCatalog []maintenance.GuestRef) ([]maintenance.DriftItem, error) {
	var vmids []int
	for _, g := range inCatalog {
		if g.Resources != nil && hasGuest(inProxmox, g.VMID) {
			vmids = append(vmids, g.VMID)
		}
	}
	if len(vmids) == 0 {
		return nil, nil
	}
	actual, err := d.resources.GuestResources(ctx, vmids)
	if err != nil {
		return nil, fmt.Errorf("read guest resources: %w", err)
	}

	var items []maintenance.DriftItem
	for _, g := range inCatalog {
		have, ok := actual[g.VMID]
		if !ok || g.Resources == nil {
			continue
		}
		if diffs := g.Resources.Differences(have); len(diffs) > 0 {
			items = append(items, maintenance.DriftItem{Kind: maintenance.ResourceMismatch, Guest: g, Differences: diffs})
		}
	}
	return items, nil
}
