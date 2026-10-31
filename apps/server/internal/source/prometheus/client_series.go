package prometheus

import (
	"context"
	"regexp"
	"strings"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

// keys are the labels the resolver may use for the client.
func (c *Client) ClientSeries(ctx context.Context, keys []string, names []network.ClientSeries, w domain.Window) ([]domain.Series, error) {
	quoted := make([]string, 0, len(keys))
	for _, k := range keys {
		quoted = append(quoted, regexp.QuoteMeta(k))
	}
	client := "client=~" + quote(strings.Join(quoted, "|"))
	rate := rateRange(w)
	queries := make([]query, 0, len(names))
	for _, n := range names {
		q := query{name: string(n), unit: domain.UnitPerSecond}
		switch n {
		case network.ClientQueriesPerSecond:
			q.expr = "sum(rate(blocky_query_total{" + client + "}" + rate + "))"
		case network.ClientBlockedPerSecond:
			q.expr = `sum(rate(blocky_client_response_total{` + client + `,response_type="BLOCKED"}` + rate + `))`
		default:
			continue
		}
		queries = append(queries, q)
	}
	return c.runQueries(ctx, queries, w)
}
