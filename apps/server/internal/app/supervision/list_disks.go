package supervision

import (
	"context"
	"fmt"

	"jarvis/server/internal/domain/supervision"
)

func (h *Hardware) Disks(ctx context.Context, node string) ([]supervision.Disk, error) {
	if _, err := h.inventoryOf(ctx, node); err != nil {
		return nil, fmt.Errorf("list disks: %w", err)
	}
	disks, err := h.source.Disks(ctx, node)
	if err != nil {
		return nil, fmt.Errorf("list disks: %w", err)
	}
	return disks, nil
}
