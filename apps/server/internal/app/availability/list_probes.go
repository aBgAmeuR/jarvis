package availability

import (
	"context"
	"fmt"

	"jarvis/server/internal/domain/availability"
)

func (p *Probes) List(ctx context.Context) ([]availability.Probe, error) {
	probes, err := p.source.Probes(ctx)
	if err != nil {
		return nil, fmt.Errorf("list probes: %w", err)
	}
	return probes, nil
}
