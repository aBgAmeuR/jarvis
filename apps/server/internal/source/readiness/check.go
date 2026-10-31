package readiness

import (
	"context"

	"jarvis/server/internal/platform/httpclient"
	"jarvis/server/internal/source/failure"
)

type Check struct {
	name    string
	path    string
	fetcher *httpclient.Fetcher
}

func New(name, path string, fetcher *httpclient.Fetcher) *Check {
	return &Check{name: name, path: path, fetcher: fetcher}
}

func (c *Check) Name() string { return c.name }

func (c *Check) Ping(ctx context.Context) error {
	return failure.Wrap(c.name, c.fetcher.Ping(ctx, c.path))
}
