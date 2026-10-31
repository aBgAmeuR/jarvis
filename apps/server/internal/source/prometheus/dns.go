package prometheus

import (
	"context"
	"sync"

	"jarvis/server/internal/domain/network"
	"jarvis/server/internal/source/failure"
)

func (c *Client) DNS(ctx context.Context) (network.DNS, error) {
	exprs := [...]string{
		"sum(increase(blocky_query_total[24h]))",
		`sum(increase(blocky_response_total{response_type="BLOCKED"}[24h]))`,
		`sum(increase(blocky_response_total{response_type="CACHED"}[24h]))`,
		`sum(rate(blocky_request_duration_seconds_sum{response_type="RESOLVED"}[24h])) / sum(rate(blocky_request_duration_seconds_count{response_type="RESOLVED"}[24h]))`,
		"sum(rate(blocky_query_total[5m]))",
	}
	var values [len(exprs)]float64
	var errs [len(exprs)]error
	var wg sync.WaitGroup
	for i, expr := range exprs {
		wg.Go(func() {
			entries, err := c.queryVector(ctx, expr)
			errs[i] = err
			if len(entries) > 0 {
				values[i] = entries[0].value
			}
		})
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return network.DNS{}, failure.Wrap(name, err)
		}
	}
	return network.DNS{
		Queries:                int64(values[0]),
		Blocked:                int64(values[1]),
		Cached:                 int64(values[2]),
		UpstreamLatencySeconds: values[3],
		QueriesPerSecond:       values[4],
	}, nil
}
