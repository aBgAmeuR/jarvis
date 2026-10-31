package loki

import (
	"context"

	"jarvis/server/internal/platform/httpclient"
	"jarvis/server/internal/source/failure"
)

const name = "loki"

type Client struct {
	fetcher *httpclient.Fetcher
}

func New(fetcher *httpclient.Fetcher) *Client {
	return &Client{fetcher: fetcher}
}

func (c *Client) Name() string { return name }

func (c *Client) Ping(ctx context.Context) error {
	return failure.Wrap(name, c.fetcher.Ping(ctx, "/ready"))
}
