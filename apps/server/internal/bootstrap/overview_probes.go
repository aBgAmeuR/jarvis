package bootstrap

import (
	"context"

	"jarvis/server/internal/app/availability"
	"jarvis/server/internal/domain/overview"
)

type overviewProbes struct {
	probes *availability.Probes
}

func (o overviewProbes) ProbeCounts(ctx context.Context) (overview.Probes, error) {
	list, err := o.probes.List(ctx)
	if err != nil {
		return overview.Probes{}, err
	}
	counts := overview.Probes{Total: len(list)}
	for _, p := range list {
		if p.Up {
			counts.Up++
		}
	}
	return counts, nil
}
