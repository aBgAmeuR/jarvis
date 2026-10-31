package supervision

import (
	"context"
	"fmt"

	"jarvis/server/internal/domain/supervision"
)

func (n *Nodes) List(ctx context.Context) ([]supervision.Node, error) {
	inventory, err := n.inventory.Inventory(ctx)
	if err != nil {
		return nil, fmt.Errorf("list nodes: %w", err)
	}
	return inventory.Nodes, nil
}
