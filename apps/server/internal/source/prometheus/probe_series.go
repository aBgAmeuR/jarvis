package prometheus

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
)

func (c *Client) ProbeSeries(ctx context.Context, probe availability.Probe, names []availability.ProbeSeries, w domain.Window) ([]domain.Series, error) {
	selector := "{job=" + quote("probe_"+string(probe.Kind)) + ",target=" + quote(probe.Target) + "}"
	rate := rateRange(w)
	queries := make([]query, 0, len(names))
	for _, n := range names {
		q := query{name: string(n)}
		switch n {
		case availability.ProbeDuration:
			q.unit, q.expr = domain.UnitSeconds, "max(probe_duration_seconds"+selector+")"
		case availability.ProbeDurationAvg:
			q.unit, q.expr = domain.UnitSeconds, "max(avg_over_time(probe_duration_seconds"+selector+rate+"))"
		case availability.ProbeDurationP99:
			q.unit, q.expr = domain.UnitSeconds, "max(quantile_over_time(0.99, probe_duration_seconds"+selector+rate+"))"
		case availability.ProbeSuccess:
			q.unit, q.expr = domain.UnitRatio, "max(probe_success"+selector+")"
		default:
			continue
		}
		queries = append(queries, q)
	}
	return c.runQueries(ctx, queries, w)
}
