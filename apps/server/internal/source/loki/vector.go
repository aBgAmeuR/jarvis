package loki

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"jarvis/server/internal/platform/httpclient"
)

type vectorResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  sampleValue       `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type vectorEntry struct {
	labels map[string]string
	value  float64
}

func (c *Client) queryVector(ctx context.Context, expr string, at time.Time) ([]vectorEntry, error) {
	params := url.Values{"query": {expr}, "time": {strconv.FormatInt(at.UnixNano(), 10)}}
	var body vectorResponse
	if err := c.fetcher.GetJSON(ctx, "/loki/api/v1/query?"+params.Encode(), &body); err != nil {
		return nil, err
	}
	if body.Status != "success" || body.Data.ResultType != "vector" {
		return nil, &httpclient.DecodeError{Err: fmt.Errorf("unexpected result %q", body.Data.ResultType)}
	}
	entries := make([]vectorEntry, 0, len(body.Data.Result))
	for _, r := range body.Data.Result {
		entries = append(entries, vectorEntry{labels: r.Metric, value: float64(r.Value)})
	}
	return entries, nil
}
