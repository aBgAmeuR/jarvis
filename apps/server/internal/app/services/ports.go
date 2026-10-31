package services

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/services"
)

type CatalogSource interface {
	ServiceCatalog(ctx context.Context) (services.Catalog, error)
}

type SignalSource interface {
	ServiceSignals(ctx context.Context) (services.Signals, error)
	ServiceSeries(ctx context.Context, name string, names []services.Series, w domain.Window) ([]domain.Series, error)
}
