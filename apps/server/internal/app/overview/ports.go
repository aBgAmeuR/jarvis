package overview

import (
	"context"

	"jarvis/server/internal/domain/overview"
)

type AlertSource interface {
	FiringAlerts(ctx context.Context) ([]overview.Alert, error)
}

type InventorySource interface {
	Inventory(ctx context.Context) (overview.Inventory, error)
}

type ProbeSource interface {
	ProbeCounts(ctx context.Context) (overview.Probes, error)
}
