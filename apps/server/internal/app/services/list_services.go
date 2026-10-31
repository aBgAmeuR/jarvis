package services

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"sync"

	"jarvis/server/internal/domain/services"
)

func (s *Services) List(ctx context.Context) ([]services.Service, error) {
	var (
		catalog services.Catalog
		signals services.Signals

		catalogErr, signalsErr error
	)
	var wg sync.WaitGroup
	wg.Go(func() { catalog, catalogErr = s.catalog.ServiceCatalog(ctx) })
	wg.Go(func() { signals, signalsErr = s.signals.ServiceSignals(ctx) })
	wg.Wait()
	for _, err := range []error{catalogErr, signalsErr} {
		if err != nil {
			return nil, fmt.Errorf("list services: %w", err)
		}
	}

	list := entries(catalog, signals)
	assignSlugs(list)
	link(list, signals)
	slices.SortFunc(list, func(a, b services.Service) int { return strings.Compare(a.Slug, b.Slug) })
	return list, nil
}
