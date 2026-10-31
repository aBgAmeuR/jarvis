package alertmanager

import (
	"context"

	"jarvis/server/internal/domain/availability"
	"jarvis/server/internal/source/failure"
)

func (c *Client) Alerts(ctx context.Context) ([]availability.Alert, error) {
	var raw []alert
	if err := c.fetcher.GetJSON(ctx, "/api/v2/alerts", &raw); err != nil {
		return nil, failure.Wrap(name, err)
	}
	alerts := make([]availability.Alert, 0, len(raw))
	for _, a := range raw {
		alerts = append(alerts, a.toDomain())
	}
	return alerts, nil
}
