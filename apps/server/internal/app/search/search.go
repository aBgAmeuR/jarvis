package search

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"strings"
	"sync"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/search"
)

const (
	DefaultLimit = 20
	MaxLimit     = 50
)

type Service struct {
	indexes []Index
}

func NewService(indexes ...Index) *Service {
	return &Service{indexes: indexes}
}

type Result struct {
	Hits     []search.Hit
	Degraded []domain.Degraded
}

// Search matches q against names, then ids and details, case-insensitively, best match first.
// An index that fails leaves its kind out and shows up in Degraded.
func (s *Service) Search(ctx context.Context, q string, limit int) (Result, error) {
	q = strings.ToLower(strings.TrimSpace(q))
	if q == "" {
		return Result{}, fmt.Errorf("%w: the query is empty", domain.ErrInvalid)
	}
	if limit == 0 {
		limit = DefaultLimit
	}
	if limit < 0 || limit > MaxLimit {
		return Result{}, fmt.Errorf("%w: limit must be between 1 and %d", domain.ErrInvalid, MaxLimit)
	}

	lists := make([][]search.Hit, len(s.indexes))
	errs := make([]error, len(s.indexes))
	var wg sync.WaitGroup
	for i, index := range s.indexes {
		wg.Go(func() { lists[i], errs[i] = index.Hits.SearchHits(ctx) })
	}
	wg.Wait()

	type scored struct {
		hit   search.Hit
		score int
	}
	var result Result
	var matches []scored
	for i, list := range lists {
		if errs[i] != nil {
			degraded := domain.NewDegraded(s.indexes[i].Source, errs[i])
			if !slices.Contains(result.Degraded, degraded) {
				result.Degraded = append(result.Degraded, degraded)
			}
			continue
		}
		for _, hit := range list {
			if score, ok := match(hit, q); ok {
				matches = append(matches, scored{hit: hit, score: score})
			}
		}
	}
	slices.SortStableFunc(matches, func(a, b scored) int {
		return cmp.Or(cmp.Compare(a.score, b.score), cmp.Compare(a.hit.Name, b.hit.Name))
	})
	for _, m := range matches[:min(len(matches), limit)] {
		result.Hits = append(result.Hits, m.hit)
	}
	return result, nil
}
