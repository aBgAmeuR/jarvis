package prometheus

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

const (
	blockyDuration = "blocky_request_duration_seconds"
	resolved       = `{response_type="RESOLVED"}`
)

func (c *Client) DNSSeries(ctx context.Context, names []network.DNSSeries, w domain.Window) ([]domain.Series, error) {
	rate := rateRange(w)
	queries := make([]query, 0, len(names))
	for _, n := range names {
		q := query{name: string(n)}
		switch n {
		case network.DNSQueriesPerSecond:
			q.unit, q.expr = domain.UnitPerSecond, "sum(rate(blocky_query_total"+rate+"))"
		case network.DNSBlockedPerSecond:
			q.unit, q.expr = domain.UnitPerSecond, `sum(rate(blocky_response_total{response_type="BLOCKED"}`+rate+`))`
		case network.DNSUpstreamLatency:
			q.unit, q.expr = domain.UnitSeconds, histogramAverage(blockyDuration, resolved, rate)
		case network.DNSUpstreamLatencyP99:
			q.unit, q.expr = domain.UnitSeconds, histogramQuantile("0.99", blockyDuration, resolved, rate)
		default:
			continue
		}
		queries = append(queries, q)
	}
	return c.runQueries(ctx, queries, w)
}
