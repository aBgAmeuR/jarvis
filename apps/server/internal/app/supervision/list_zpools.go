package supervision

import (
	"context"
	"fmt"

	"jarvis/server/internal/domain/supervision"
)

func (h *Hardware) ZPools(ctx context.Context, node string) ([]supervision.ZPool, error) {
	if _, err := h.inventoryOf(ctx, node); err != nil {
		return nil, fmt.Errorf("list zpools: %w", err)
	}
	pools, err := h.source.ZPools(ctx, node)
	if err != nil {
		return nil, fmt.Errorf("list zpools: %w", err)
	}
	return pools, nil
}
