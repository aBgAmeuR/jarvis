package bootstrap

import (
	"context"
	"strconv"

	"jarvis/server/internal/app/availability"
	domainavailability "jarvis/server/internal/domain/availability"
	"jarvis/server/internal/domain/supervision"
)

// guestAlerts lets supervision read alerts without importing availability.
type guestAlerts struct {
	alerts *availability.Alerts
}

func (g guestAlerts) GuestAlerts(ctx context.Context) ([]supervision.GuestAlert, error) {
	firing, err := g.alerts.List(ctx, availability.AlertFilter{State: domainavailability.AlertFiring})
	if err != nil {
		return nil, err
	}
	out := make([]supervision.GuestAlert, 0, len(firing))
	for _, alert := range firing {
		vmid, err := strconv.Atoi(alert.Labels["vmid"])
		if err != nil {
			continue
		}
		out = append(out, supervision.GuestAlert{VMID: vmid, Severity: alert.Severity})
	}
	return out, nil
}
