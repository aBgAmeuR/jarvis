package supervision

import (
	"context"
	"fmt"
	"slices"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/supervision"
)

func (m *Metrics) Node(ctx context.Context, node string, names []supervision.NodeSeries, start, end time.Time) (domain.TimeSeries, error) {
	w, err := domain.NewWindow(start, end, m.now())
	if err != nil {
		return domain.TimeSeries{}, err
	}
	inventory, err := m.inventory.Inventory(ctx)
	if err != nil {
		return domain.TimeSeries{}, fmt.Errorf("node metrics: %w", err)
	}
	if !slices.ContainsFunc(inventory.Nodes, func(n supervision.Node) bool { return n.Name == node }) {
		return domain.TimeSeries{}, domain.NewNotFound("node", node)
	}
	if len(names) == 0 {
		names = supervision.AllNodeSeries()
	}
	series, err := m.series.NodeSeries(ctx, node, slices.Compact(slices.Sorted(slices.Values(names))), w)
	if err != nil {
		return domain.TimeSeries{}, fmt.Errorf("node %s metrics: %w", node, err)
	}
	return domain.TimeSeries{Window: w, Series: series}, nil
}
