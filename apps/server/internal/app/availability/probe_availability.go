package availability

import (
	"context"
	"fmt"

	"jarvis/server/internal/domain/availability"
)

// Availability covers the last 90 UTC days and today so far.
func (p *Probes) Availability(ctx context.Context, id string) ([]availability.Bucket, error) {
	probe, err := p.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	buckets, err := p.source.ProbeAvailability(ctx, probe, p.now())
	if err != nil {
		return nil, fmt.Errorf("probe %s availability: %w", id, err)
	}
	return buckets, nil
}
