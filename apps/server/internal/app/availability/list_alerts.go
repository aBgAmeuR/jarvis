package availability

import (
	"cmp"
	"context"
	"fmt"
	"slices"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
)

// Alerts without a known severity, like Watchdog, are left out.
func (a *Alerts) List(ctx context.Context, f AlertFilter) ([]availability.Alert, error) {
	all, err := a.source.Alerts(ctx)
	if err != nil {
		return nil, fmt.Errorf("list alerts: %w", err)
	}
	alerts := slices.Clone(all)
	alerts = slices.DeleteFunc(alerts, func(alert availability.Alert) bool {
		return alert.Severity == domain.SeverityNone || !f.matches(alert)
	})
	slices.SortFunc(alerts, func(x, y availability.Alert) int {
		return cmp.Or(
			cmp.Compare(x.Severity.Rank(), y.Severity.Rank()),
			y.StartsAt.Compare(x.StartsAt),
			cmp.Compare(x.Fingerprint, y.Fingerprint),
		)
	})
	return alerts, nil
}
