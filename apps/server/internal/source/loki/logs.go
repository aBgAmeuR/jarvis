package loki

import (
	"cmp"
	"context"
	"errors"
	"net/url"
	"slices"
	"strconv"
	"time"

	"jarvis/server/internal/domain/logs"
	"jarvis/server/internal/platform/httpclient"
	"jarvis/server/internal/source/failure"
)

type streamsResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Stream map[string]string `json:"stream"`
			Values [][2]string       `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

func (c *Client) Logs(ctx context.Context, q logs.Query) ([]logs.Entry, error) {
	params := url.Values{
		"query":     {logQL(q)},
		"start":     {strconv.FormatInt(q.Start.UnixNano(), 10)},
		"end":       {strconv.FormatInt(q.End.UnixNano(), 10)},
		"limit":     {strconv.Itoa(q.Limit)},
		"direction": {"backward"},
	}
	var body streamsResponse
	if err := c.fetcher.GetJSON(ctx, "/loki/api/v1/query_range?"+params.Encode(), &body); err != nil {
		return nil, failure.Wrap(name, err)
	}
	if body.Status != "success" || body.Data.ResultType != "streams" {
		return nil, failure.Wrap(name, &httpclient.DecodeError{Err: errUnexpectedResult})
	}

	entries := []logs.Entry{}
	for _, stream := range body.Data.Result {
		vmid, _ := strconv.Atoi(stream.Stream["vmid"])
		for _, value := range stream.Values {
			nanos, err := strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				return nil, failure.Wrap(name, &httpclient.DecodeError{Err: err})
			}
			entries = append(entries, logs.Entry{
				Timestamp: time.Unix(0, nanos).UTC(),
				Line:      value[1],
				Level:     stream.Stream["level"],
				Service:   stream.Stream["service_name"],
				Instance:  stream.Stream["instance"],
				VMID:      vmid,
			})
		}
	}
	slices.SortFunc(entries, func(a, b logs.Entry) int { return cmp.Compare(b.Timestamp.UnixNano(), a.Timestamp.UnixNano()) })
	return entries[:min(len(entries), q.Limit)], nil
}

var errUnexpectedResult = errors.New("unexpected result type")
