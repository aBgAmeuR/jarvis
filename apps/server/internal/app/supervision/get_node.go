package supervision

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/supervision"
)

func (n *Nodes) Get(ctx context.Context, name string) (supervision.Node, error) {
	nodes, err := n.List(ctx)
	if err != nil {
		return supervision.Node{}, err
	}
	for _, node := range nodes {
		if node.Name == name {
			return node, nil
		}
	}
	return supervision.Node{}, domain.NewNotFound("node", name)
}
