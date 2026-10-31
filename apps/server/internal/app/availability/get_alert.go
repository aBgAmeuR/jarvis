package availability

import (
	"context"
	"fmt"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
)

func (a *Alerts) Get(ctx context.Context, fingerprint string) (availability.Alert, error) {
	alerts, err := a.source.Alerts(ctx)
	if err != nil {
		return availability.Alert{}, fmt.Errorf("get alert: %w", err)
	}
	for _, alert := range alerts {
		if alert.Fingerprint == fingerprint {
			return alert, nil
		}
	}
	return availability.Alert{}, domain.NewNotFound("alert", fingerprint)
}
