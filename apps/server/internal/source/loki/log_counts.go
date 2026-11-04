package loki

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/logs"
	"jarvis/server/internal/platform/httpclient"
	"jarvis/server/internal/source/failure"
)

type matrixResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Values []samplePoint     `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

func (c *Client) LogCounts(ctx context.Context, q logs.Query, w domain.Window) ([]logs.Count, error) {
	step := strconv.FormatInt(int64(w.Step/time.Second), 10) + "s"
	params := url.Values{
		"query": {"sum by (level) (count_over_time(" + logQL(q) + " [" + step + "]))"},
		"start": {strconv.FormatInt(w.Start.UnixNano(), 10)},
		"end":   {strconv.FormatInt(w.End.UnixNano(), 10)},
		"step":  {step},
	}
	var body matrixResponse
	if err := c.fetcher.GetJSON(ctx, "/loki/api/v1/query_range?"+params.Encode(), &body); err != nil {
		return nil, failure.Wrap(name, err)
	}
	if body.Status != "success" || body.Data.ResultType != "matrix" {
		return nil, failure.Wrap(name, &httpclient.DecodeError{Err: errUnexpectedResult})
	}

	counts := []logs.Count{}
	for _, series := range body.Data.Result {
		for _, point := range series.Values {
			if point.value > 0 {
				counts = append(counts, logs.Count{Time: point.time, Level: series.Metric["level"], N: int(point.value)})
			}
		}
	}
	return counts, nil
}
