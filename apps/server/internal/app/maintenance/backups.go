package maintenance

import (
	"cmp"
	"context"
	"fmt"
	"slices"

	"jarvis/server/internal/domain/maintenance"
)

type Backups struct {
	guests GuestSource
	source BackupSource
}

func NewBackups(guests GuestSource, source BackupSource) *Backups {
	return &Backups{guests: guests, source: source}
}

func (b *Backups) Jobs(ctx context.Context) ([]maintenance.BackupJob, error) {
	jobs, err := b.source.BackupJobs(ctx)
	if err != nil {
		return nil, fmt.Errorf("list backup jobs: %w", err)
	}
	jobs = slices.Clone(jobs)
	slices.SortFunc(jobs, func(x, y maintenance.BackupJob) int { return cmp.Compare(x.ID, y.ID) })
	return jobs, nil
}
