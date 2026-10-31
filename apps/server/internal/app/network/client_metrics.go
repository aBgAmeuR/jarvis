package network

import (
	"context"
	"fmt"
	"slices"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

func (c *Clients) Metrics(ctx context.Context, id string, names []network.ClientSeries, start, end time.Time) (domain.TimeSeries, error) {
	w, err := domain.NewWindow(start, end, c.now())
	if err != nil {
		return domain.TimeSeries{}, err
	}
	client, _, err := c.Get(ctx, id)
	if err != nil {
		return domain.TimeSeries{}, err
	}
	if len(names) == 0 {
		names = network.AllClientSeries()
	}
	series, err := c.source.ClientSeries(ctx, labelsOf(client.IP, client.Name), slices.Compact(slices.Sorted(slices.Values(names))), w)
	if err != nil {
		return domain.TimeSeries{}, fmt.Errorf("client %s metrics: %w", id, err)
	}
	return domain.TimeSeries{Window: w, Series: series}, nil
}
