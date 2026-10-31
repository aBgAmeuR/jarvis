package bootstrap

import (
	appavailability "jarvis/server/internal/app/availability"
	appsupervision "jarvis/server/internal/app/supervision"
	httpsupervision "jarvis/server/internal/httpapi/supervision"
)

func newSupervision(s sources, alerts *appavailability.Alerts) httpsupervision.Services {
	return httpsupervision.Services{
		Guests:   appsupervision.NewGuests(s.proxmox, s.catalog, guestAlerts{alerts: alerts}, s.prometheus),
		Nodes:    appsupervision.NewNodes(s.proxmox),
		Metrics:  appsupervision.NewMetrics(s.proxmox, s.prometheus),
		Hardware: appsupervision.NewHardware(s.proxmox, s.prometheus),
	}
}
