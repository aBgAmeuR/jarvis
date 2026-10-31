package bootstrap

import (
	appavailability "jarvis/server/internal/app/availability"
	httpavailability "jarvis/server/internal/httpapi/availability"
)

func newAvailability(s sources) httpavailability.Services {
	return httpavailability.Services{
		Alerts:       appavailability.NewAlerts(s.alertmanager, s.prometheus),
		Probes:       appavailability.NewProbes(s.prometheus),
		Certificates: appavailability.NewCertificates(s.prometheus),
		Silences:     appavailability.NewSilences(s.alertmanager),
	}
}
