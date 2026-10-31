package services

import (
	"context"
	"fmt"
	"slices"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/services"
)

func (s *Services) Metrics(ctx context.Context, slug string, names []services.Series, start, end time.Time) (domain.TimeSeries, error) {
	w, err := domain.NewWindow(start, end, s.now())
	if err != nil {
		return domain.TimeSeries{}, err
	}
	service, err := s.Get(ctx, slug)
	if err != nil {
		return domain.TimeSeries{}, err
	}
	if len(names) == 0 {
		names = services.AllSeries()
	}
	series, err := s.signals.ServiceSeries(ctx, service.Name, slices.Compact(slices.Sorted(slices.Values(names))), w)
	if err != nil {
		return domain.TimeSeries{}, fmt.Errorf("service %s metrics: %w", slug, err)
	}
	return domain.TimeSeries{Window: w, Series: series}, nil
}
