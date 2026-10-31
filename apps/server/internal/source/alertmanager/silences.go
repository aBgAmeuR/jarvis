package alertmanager

import (
	"context"

	"jarvis/server/internal/domain/availability"
	"jarvis/server/internal/source/failure"
)

func (c *Client) Silences(ctx context.Context) ([]availability.Silence, error) {
	var raw []silence
	if err := c.fetcher.GetJSON(ctx, "/api/v2/silences", &raw); err != nil {
		return nil, failure.Wrap(name, err)
	}
	silences := make([]availability.Silence, 0, len(raw))
	for _, s := range raw {
		silences = append(silences, s.toDomain())
	}
	return silences, nil
}
