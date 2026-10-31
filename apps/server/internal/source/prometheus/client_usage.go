package prometheus

import (
	"context"
	"sync"
	"time"

	"jarvis/server/internal/domain/network"
	"jarvis/server/internal/source/failure"
)

// The last query is the last 5-minute step in which the counter of the client moved.
func (c *Client) ClientUsage(ctx context.Context) (map[string]network.Usage, error) {
	exprs := [...]string{
		"sum by (client) (increase(blocky_query_total[24h]))",
		`sum by (client) (increase(blocky_client_response_total{response_type="BLOCKED"}[24h]))`,
		"max by (client) (max_over_time(timestamp(changes(blocky_query_total[5m]) > 0)[24h:5m]))",
	}
	var results [len(exprs)][]entry
	var errs [len(exprs)]error
	var wg sync.WaitGroup
	for i, expr := range exprs {
		wg.Go(func() { results[i], errs[i] = c.queryVector(ctx, expr) })
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return nil, failure.Wrap(name, err)
		}
	}

	usage := map[string]network.Usage{}
	for _, e := range results[0] {
		u := usage[e.labels["client"]]
		u.Queries = e.value
		usage[e.labels["client"]] = u
	}
	for _, e := range results[1] {
		u := usage[e.labels["client"]]
		u.Blocked = e.value
		usage[e.labels["client"]] = u
	}
	for _, e := range results[2] {
		u := usage[e.labels["client"]]
		at := time.Unix(int64(e.value), 0).UTC()
		u.LastQueryAt = &at
		usage[e.labels["client"]] = u
	}
	return usage, nil
}
