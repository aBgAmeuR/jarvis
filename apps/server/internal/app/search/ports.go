package search

import (
	"context"

	"jarvis/server/internal/domain/search"
)

// HitSource lists every object of one kind, unfiltered.
type HitSource interface {
	SearchHits(ctx context.Context) ([]search.Hit, error)
}

// Index is one kind of object. Source names the data source reported when the error names none.
type Index struct {
	Source string
	Hits   HitSource
}
