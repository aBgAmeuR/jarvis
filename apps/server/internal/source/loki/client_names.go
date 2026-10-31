package loki

import (
	"context"
	"strconv"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/source/failure"
)

func (c *Client) ClientNames(ctx context.Context, w domain.Window) (map[string]string, error) {
	expr := `sum by (client_ip, client_names) (count_over_time({service_name="blocky"} | prefix="queryLog" [` + strconv.FormatInt(int64(w.End.Sub(w.Start).Seconds()), 10) + `s]))`
	entries, err := c.queryVector(ctx, expr, w.End)
	if err != nil {
		return nil, failure.Wrap(name, err)
	}
	names := make(map[string]string, len(entries))
	for _, e := range entries {
		names[e.labels["client_ip"]] = e.labels["client_names"]
	}
	return names, nil
}
