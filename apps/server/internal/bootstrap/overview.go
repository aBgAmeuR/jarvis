package bootstrap

import (
	appavailability "jarvis/server/internal/app/availability"
	appoverview "jarvis/server/internal/app/overview"
)

func newOverview(s sources, alerts *appavailability.Alerts, probes *appavailability.Probes) *appoverview.Service {
	return appoverview.NewService(overviewAlerts{alerts: alerts}, overviewInventory{proxmox: s.proxmox}, overviewProbes{probes: probes})
}
