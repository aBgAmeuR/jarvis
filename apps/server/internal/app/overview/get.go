package overview

import (
	"context"
	"sync"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/overview"
)

type Service struct {
	alerts    AlertSource
	inventory InventorySource
	probes    ProbeSource
}

func NewService(alerts AlertSource, inventory InventorySource, probes ProbeSource) *Service {
	return &Service{alerts: alerts, inventory: inventory, probes: probes}
}

type Result struct {
	Overview overview.Overview
	Degraded []domain.Degraded
}

// Get never fails: a source that does not answer leaves its part out and shows up in Degraded.
func (s *Service) Get(ctx context.Context) Result {
	var (
		alerts    []overview.Alert
		inventory overview.Inventory
		probes    overview.Probes

		alertsErr, inventoryErr, probesErr error
	)
	var wg sync.WaitGroup
	wg.Go(func() { alerts, alertsErr = s.alerts.FiringAlerts(ctx) })
	wg.Go(func() { inventory, inventoryErr = s.inventory.Inventory(ctx) })
	wg.Go(func() { probes, probesErr = s.probes.ProbeCounts(ctx) })
	wg.Wait()

	var result Result
	if alertsErr != nil {
		result.Degraded = append(result.Degraded, domain.NewDegraded("alertmanager", alertsErr))
		result.Overview.Status = overview.StatusUnknown
	} else {
		result.Overview.Status = statusOf(alerts)
		result.Overview.ToHandle = alerts
	}
	if inventoryErr != nil {
		result.Degraded = append(result.Degraded, domain.NewDegraded("proxmox", inventoryErr))
	} else {
		result.Overview.Nodes = inventory.Nodes
		result.Overview.Guests = &inventory.Guests
	}
	if probesErr != nil {
		result.Degraded = append(result.Degraded, domain.NewDegraded("prometheus", probesErr))
	} else {
		result.Overview.Probes = &probes
	}
	return result
}
