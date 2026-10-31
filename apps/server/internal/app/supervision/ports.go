package supervision

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/supervision"
)

type InventorySource interface {
	Inventory(ctx context.Context) (supervision.Inventory, error)
}

type CatalogSource interface {
	Catalog(ctx context.Context) (supervision.Catalog, error)
}

type MetricsSource interface {
	GuestSeries(ctx context.Context, vmid int, names []supervision.GuestSeries, w domain.Window) ([]domain.Series, error)
	NodeSeries(ctx context.Context, node string, names []supervision.NodeSeries, w domain.Window) ([]domain.Series, error)
}

type HardwareSource interface {
	ZPools(ctx context.Context, node string) ([]supervision.ZPool, error)
	Disks(ctx context.Context, node string) ([]supervision.Disk, error)
}

// UpgradeSource reports the pending apt upgrades per guest VMID.
type UpgradeSource interface {
	Upgrades(ctx context.Context) (map[int]supervision.Upgrades, error)
}

type GuestAlertSource interface {
	GuestAlerts(ctx context.Context) ([]supervision.GuestAlert, error)
}
