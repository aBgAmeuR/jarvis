package availability

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
)

func (a *Alerts) Resolved(ctx context.Context, f AlertFilter, start, end time.Time) ([]availability.Alert, error) {
	w, err := domain.NewWindow(start, end, a.now())
	if err != nil {
		return nil, err
	}
	all, err := a.history.ResolvedAlerts(ctx, w)
	if err != nil {
		return nil, fmt.Errorf("list resolved alerts: %w", err)
	}
	f.State = availability.AlertResolved
	alerts := slices.DeleteFunc(slices.Clone(all), func(alert availability.Alert) bool {
		return alert.Severity == domain.SeverityNone || !f.matches(alert)
	})
	slices.SortFunc(alerts, func(x, y availability.Alert) int {
		return cmp.Or(y.EndsAt.Compare(*x.EndsAt), cmp.Compare(x.Fingerprint, y.Fingerprint))
	})
	return alerts, nil
}
