package prometheus

import (
	"context"
	"fmt"
	"net/url"

	"jarvis/server/internal/platform/httpclient"
)

type vectorResponse struct {
	Status    string `json:"status"`
	ErrorType string `json:"errorType"`
	Data      struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  sample            `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type entry struct {
	labels map[string]string
	value  float64
}

func (c *Client) queryVector(ctx context.Context, expr string) ([]entry, error) {
	var body vectorResponse
	if err := c.fetcher.GetJSON(ctx, "/api/v1/query?"+url.Values{"query": {expr}}.Encode(), &body); err != nil {
		return nil, err
	}
	if body.Status != "success" || body.Data.ResultType != "vector" {
		return nil, &httpclient.DecodeError{Err: fmt.Errorf("query failed: %s %s", body.Status, body.ErrorType)}
	}
	entries := make([]entry, 0, len(body.Data.Result))
	for _, r := range body.Data.Result {
		if r.Value.finite() {
			entries = append(entries, entry{labels: r.Metric, value: r.Value.value})
		}
	}
	return entries, nil
}
