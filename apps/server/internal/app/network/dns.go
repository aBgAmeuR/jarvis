package network

import (
	"context"
	"fmt"
	"slices"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

type DNS struct {
	source  DNSSource
	domains DomainSource
	now     func() time.Time
}

func NewDNS(source DNSSource, domains DomainSource) *DNS {
	return &DNS{source: source, domains: domains, now: time.Now}
}

func (d *DNS) Summary(ctx context.Context) (network.DNS, error) {
	summary, err := d.source.DNS(ctx)
	if err != nil {
		return network.DNS{}, fmt.Errorf("dns summary: %w", err)
	}
	return summary, nil
}

func (d *DNS) Metrics(ctx context.Context, names []network.DNSSeries, start, end time.Time) (domain.TimeSeries, error) {
	w, err := domain.NewWindow(start, end, d.now())
	if err != nil {
		return domain.TimeSeries{}, err
	}
	if len(names) == 0 {
		names = network.AllDNSSeries()
	}
	series, err := d.source.DNSSeries(ctx, slices.Compact(slices.Sorted(slices.Values(names))), w)
	if err != nil {
		return domain.TimeSeries{}, fmt.Errorf("dns metrics: %w", err)
	}
	return domain.TimeSeries{Window: w, Series: series}, nil
}
