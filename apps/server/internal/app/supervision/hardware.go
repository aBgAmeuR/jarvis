package supervision

import (
	"context"
	"slices"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/supervision"
)

type Hardware struct {
	inventory InventorySource
	source    HardwareSource
}

func NewHardware(inventory InventorySource, source HardwareSource) *Hardware {
	return &Hardware{inventory: inventory, source: source}
}

func (h *Hardware) inventoryOf(ctx context.Context, node string) (supervision.Inventory, error) {
	inventory, err := h.inventory.Inventory(ctx)
	if err != nil {
		return supervision.Inventory{}, err
	}
	if !slices.ContainsFunc(inventory.Nodes, func(n supervision.Node) bool { return n.Name == node }) {
		return supervision.Inventory{}, domain.NewNotFound("node", node)
	}
	return inventory, nil
}
