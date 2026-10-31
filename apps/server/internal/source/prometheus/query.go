package prometheus

import (
	"context"
	"sync"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/source/failure"
)

type query struct {
	name string
	unit domain.Unit
	expr string
}

func (c *Client) runQueries(ctx context.Context, queries []query, w domain.Window) ([]domain.Series, error) {
	series := make([]domain.Series, len(queries))
	errs := make([]error, len(queries))
	var wg sync.WaitGroup
	for i, q := range queries {
		wg.Go(func() { series[i], errs[i] = c.queryRange(ctx, q.expr, q.unit, q.name, w) })
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return nil, failure.Wrap(name, err)
		}
	}
	return series, nil
}
