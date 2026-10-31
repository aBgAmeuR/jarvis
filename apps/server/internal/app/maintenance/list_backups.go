package maintenance

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"sync"

	"jarvis/server/internal/domain/maintenance"
)

func (b *Backups) List(ctx context.Context, guest int) ([]maintenance.Backup, error) {
	var (
		guests    []maintenance.GuestRef
		uncovered []int
		volumes   []maintenance.Volume
		tasks     []maintenance.Task

		errs [4]error
	)
	var wg sync.WaitGroup
	wg.Go(func() { guests, errs[0] = b.guests.BackupGuests(ctx) })
	wg.Go(func() { uncovered, errs[1] = b.source.UncoveredGuests(ctx) })
	wg.Go(func() { volumes, errs[2] = b.source.BackupVolumes(ctx) })
	wg.Go(func() { tasks, errs[3] = b.source.BackupTasks(ctx) })
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return nil, fmt.Errorf("list backups: %w", err)
		}
	}

	backups := make([]maintenance.Backup, 0, len(guests))
	for _, g := range guests {
		if guest != 0 && g.VMID != guest {
			continue
		}
		backups = append(backups, summarize(g, !slices.Contains(uncovered, g.VMID), volumes, tasks))
	}
	slices.SortFunc(backups, func(x, y maintenance.Backup) int { return cmp.Compare(x.Guest.VMID, y.Guest.VMID) })
	return backups, nil
}
