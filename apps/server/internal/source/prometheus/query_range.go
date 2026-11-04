package prometheus

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/platform/httpclient"
)

type matrixResponse struct {
	Status    string `json:"status"`
	ErrorType string `json:"errorType"`
	Data      struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Values []sample          `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

type matrixSeries struct {
	labels  map[string]string
	samples []sample
}

func (c *Client) queryMatrix(ctx context.Context, expr string, w domain.Window) ([]matrixSeries, error) {
	return c.queryMatrixBetween(ctx, expr, w.Start, w.End, w.Step)
}

func (c *Client) queryMatrixBetween(ctx context.Context, expr string, start, end time.Time, step time.Duration) ([]matrixSeries, error) {
	params := url.Values{
		"query": {expr},
		"start": {strconv.FormatInt(start.Unix(), 10)},
		"end":   {strconv.FormatInt(end.Unix(), 10)},
		"step":  {strconv.FormatInt(int64(step.Seconds()), 10)},
	}
	var body matrixResponse
	if err := c.fetcher.GetJSON(ctx, "/api/v1/query_range?"+params.Encode(), &body); err != nil {
		return nil, err
	}
	if body.Status != "success" || body.Data.ResultType != "matrix" {
		return nil, &httpclient.DecodeError{Err: fmt.Errorf("query failed: %s %s", body.Status, body.ErrorType)}
	}
	out := make([]matrixSeries, 0, len(body.Data.Result))
	for _, r := range body.Data.Result {
		series := matrixSeries{labels: r.Metric, samples: make([]sample, 0, len(r.Values))}
		for _, s := range r.Values {
			if s.finite() {
				series.samples = append(series.samples, s)
			}
		}
		out = append(out, series)
	}
	return out, nil
}

// expr must yield a single series; an absent one gives an empty result.
func (c *Client) queryRange(ctx context.Context, expr string, unit domain.Unit, seriesName string, w domain.Window) (domain.Series, error) {
	matrix, err := c.queryMatrix(ctx, expr, w)
	if err != nil {
		return domain.Series{}, err
	}
	series := domain.Series{Name: seriesName, Unit: unit, Query: expr, Timestamps: []time.Time{}, Values: []float64{}}
	if len(matrix) > 0 {
		for _, s := range matrix[0].samples {
			series.Timestamps = append(series.Timestamps, s.at)
			series.Values = append(series.Values, s.value)
		}
	}
	return series, nil
}
