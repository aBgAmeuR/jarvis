package supervision

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"jarvis/server/internal/domain/supervision"
)

func (h *Hardware) Storages(ctx context.Context, node string) ([]supervision.Storage, error) {
	inventory, err := h.inventoryOf(ctx, node)
	if err != nil {
		return nil, fmt.Errorf("list storages: %w", err)
	}
	storages := slices.DeleteFunc(slices.Clone(inventory.Storages), func(s supervision.Storage) bool { return s.Node != node })
	slices.SortFunc(storages, func(a, b supervision.Storage) int { return strings.Compare(a.Name, b.Name) })
	return storages, nil
}
