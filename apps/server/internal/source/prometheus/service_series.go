package prometheus

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/services"
)

func (c *Client) ServiceSeries(ctx context.Context, service string, names []services.Series, w domain.Window) ([]domain.Series, error) {
	selector := "{service_name=" + quote(service) + "}"
	failed := `{service_name=` + quote(service) + `,http_response_status_code=~"5.."}`
	rate := rateRange(w)
	queries := make([]query, 0, len(names))
	for _, n := range names {
		q := query{name: string(n)}
		switch n {
		case services.RequestsPerSecond:
			q.unit, q.expr = domain.UnitPerSecond, "sum(rate(http_server_request_duration_seconds_count"+selector+rate+"))"
		case services.ErrorRatio:
			q.unit = domain.UnitRatio
			q.expr = "(sum(rate(http_server_request_duration_seconds_count" + failed + rate + ")) or vector(0)) / sum(rate(http_server_request_duration_seconds_count" + selector + rate + "))"
		case services.LatencyP95Seconds:
			q.unit = domain.UnitSeconds
			q.expr = "histogram_quantile(0.95, sum by (le) (rate(http_server_request_duration_seconds_bucket" + selector + rate + ")))"
		default:
			continue
		}
		queries = append(queries, q)
	}
	return c.runQueries(ctx, queries, w)
}
