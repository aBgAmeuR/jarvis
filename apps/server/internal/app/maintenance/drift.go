package maintenance

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"sync"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/maintenance"
)

type Drift struct {
	proxmox   GuestSource
	catalog   CatalogSource
	resources ResourceSource
}

func NewDrift(proxmox GuestSource, catalog CatalogSource, resources ResourceSource) *Drift {
	return &Drift{proxmox: proxmox, catalog: catalog, resources: resources}
}

type DriftList struct {
	Items    []maintenance.DriftItem
	Degraded []domain.Degraded
}

// The guest lists are required: failing to read the resources only degrades the result.
func (d *Drift) List(ctx context.Context) (DriftList, error) {
	var (
		inProxmox, inCatalog []maintenance.GuestRef
		proxmoxErr           error
		catalogErr           error
	)
	var wg sync.WaitGroup
	wg.Go(func() { inProxmox, proxmoxErr = d.proxmox.BackupGuests(ctx) })
	wg.Go(func() { inCatalog, catalogErr = d.catalog.CatalogGuests(ctx) })
	wg.Wait()
	for _, err := range []error{proxmoxErr, catalogErr} {
		if err != nil {
			return DriftList{}, fmt.Errorf("list drift items: %w", err)
		}
	}

	items := missingItems(inProxmox, inCatalog)
	var degraded []domain.Degraded
	mismatches, err := d.resourceMismatches(ctx, inProxmox, inCatalog)
	if err != nil {
		degraded = append(degraded, domain.NewDegraded("proxmox", err))
	}
	items = append(items, mismatches...)
	slices.SortFunc(items, func(x, y maintenance.DriftItem) int {
		return cmp.Or(cmp.Compare(x.Kind, y.Kind), cmp.Compare(x.Guest.VMID, y.Guest.VMID))
	})
	return DriftList{Items: items, Degraded: degraded}, nil
}
