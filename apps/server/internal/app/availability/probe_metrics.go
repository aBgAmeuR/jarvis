package availability

import (
	"context"
	"fmt"
	"slices"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
)

func (p *Probes) Metrics(ctx context.Context, id string, names []availability.ProbeSeries, start, end time.Time) (domain.TimeSeries, error) {
	w, err := domain.NewWindow(start, end, p.now())
	if err != nil {
		return domain.TimeSeries{}, err
	}
	probe, err := p.Get(ctx, id)
	if err != nil {
		return domain.TimeSeries{}, err
	}
	if len(names) == 0 {
		names = availability.AllProbeSeries()
	}
	series, err := p.source.ProbeSeries(ctx, probe, slices.Compact(slices.Sorted(slices.Values(names))), w)
	if err != nil {
		return domain.TimeSeries{}, fmt.Errorf("probe %s metrics: %w", id, err)
	}
	return domain.TimeSeries{Window: w, Series: series}, nil
}
