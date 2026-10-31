package maintenance

import (
	"context"

	"jarvis/server/internal/domain/maintenance"
)

type GuestSource interface {
	BackupGuests(ctx context.Context) ([]maintenance.GuestRef, error)
}

type CatalogSource interface {
	CatalogGuests(ctx context.Context) ([]maintenance.GuestRef, error)
}

type ResourceSource interface {
	// Virtual machines are absent from the result.
	GuestResources(ctx context.Context, vmids []int) (map[int]maintenance.Resources, error)
}

type BackupSource interface {
	BackupJobs(ctx context.Context) ([]maintenance.BackupJob, error)
	UncoveredGuests(ctx context.Context) ([]int, error)
	BackupVolumes(ctx context.Context) ([]maintenance.Volume, error)
	BackupTasks(ctx context.Context) ([]maintenance.Task, error)
}
