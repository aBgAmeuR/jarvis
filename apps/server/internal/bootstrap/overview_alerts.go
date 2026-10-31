package bootstrap

import (
	"context"

	"jarvis/server/internal/app/availability"
	domainavailability "jarvis/server/internal/domain/availability"
	"jarvis/server/internal/domain/overview"
)

type overviewAlerts struct {
	alerts *availability.Alerts
}

func (o overviewAlerts) FiringAlerts(ctx context.Context) ([]overview.Alert, error) {
	firing, err := o.alerts.List(ctx, availability.AlertFilter{State: domainavailability.AlertFiring})
	if err != nil {
		return nil, err
	}
	out := make([]overview.Alert, 0, len(firing))
	for _, a := range firing {
		out = append(out, overview.Alert{
			Fingerprint: a.Fingerprint,
			Name:        a.Name,
			Severity:    a.Severity,
			StartsAt:    a.StartsAt,
			Summary:     a.Annotations["summary"],
		})
	}
	return out, nil
}
