package availability

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
)

func (p *Probes) Get(ctx context.Context, id string) (availability.Probe, error) {
	probes, err := p.List(ctx)
	if err != nil {
		return availability.Probe{}, err
	}
	for _, probe := range probes {
		if probe.ID == id {
			return probe, nil
		}
	}
	return availability.Probe{}, domain.NewNotFound("probe", id)
}
