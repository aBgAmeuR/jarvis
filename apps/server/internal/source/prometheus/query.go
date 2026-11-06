package prometheus

import (
	"context"
	"sync"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/source/failure"
)

type query struct {
	name    string
	unit    domain.Unit
	expr    string
	samples string
}

func gauge(name string, unit domain.Unit, samples string) query {
	return query{name: name, unit: unit, expr: "max(" + samples + ")", samples: samples}
}

func (c *Client) runQueries(ctx context.Context, queries []query, w domain.Window) ([]domain.Series, error) {
	series := make([]domain.Series, len(queries))
	errs := make([]error, len(queries))
	var wg sync.WaitGroup
	for i, q := range queries {
		wg.Go(func() { series[i], errs[i] = c.querySeries(ctx, q, w) })
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return nil, failure.Wrap(name, err)
		}
	}
	return series, nil
}
