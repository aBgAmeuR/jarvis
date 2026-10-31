package supervision

import (
	"context"
	"fmt"
	"sync"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/supervision"
)

type GuestList struct {
	Items    []supervision.Guest
	Degraded []domain.Degraded
}

// The inventory is required: a failing catalog or alert source only degrades the result.
func (g *Guests) List(ctx context.Context) (GuestList, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var (
		inventory supervision.Inventory
		catalog   supervision.Catalog
		alerts    []supervision.GuestAlert
		upgrades  map[int]supervision.Upgrades

		inventoryErr, catalogErr, alertsErr, upgradesErr error
	)
	var wg sync.WaitGroup
	wg.Go(func() {
		inventory, inventoryErr = g.inventory.Inventory(ctx)
		if inventoryErr != nil {
			cancel()
		}
	})
	wg.Go(func() { catalog, catalogErr = g.catalog.Catalog(ctx) })
	wg.Go(func() { alerts, alertsErr = g.alerts.GuestAlerts(ctx) })
	wg.Go(func() { upgrades, upgradesErr = g.upgrades.Upgrades(ctx) })
	wg.Wait()

	if inventoryErr != nil {
		return GuestList{}, fmt.Errorf("list guests: %w", inventoryErr)
	}

	var degraded []domain.Degraded
	var joinCatalog *supervision.Catalog
	if catalogErr != nil {
		degraded = append(degraded, domain.NewDegraded("catalog", catalogErr))
	} else {
		joinCatalog = &catalog
	}
	if alertsErr != nil {
		degraded = append(degraded, domain.NewDegraded("alertmanager", alertsErr))
	}
	if upgradesErr != nil {
		degraded = append(degraded, domain.NewDegraded("prometheus", upgradesErr))
	}
	return GuestList{Items: enrich(inventory.Guests, joinCatalog, alerts, upgrades), Degraded: degraded}, nil
}
