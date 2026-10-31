package supervision

import (
	"context"
	"fmt"
	"slices"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/supervision"
)

// A zero start or end takes the defaults of domain.NewWindow.
func (m *Metrics) Guest(ctx context.Context, vmid int, names []supervision.GuestSeries, start, end time.Time) (domain.TimeSeries, error) {
	w, err := domain.NewWindow(start, end, m.now())
	if err != nil {
		return domain.TimeSeries{}, err
	}
	inventory, err := m.inventory.Inventory(ctx)
	if err != nil {
		return domain.TimeSeries{}, fmt.Errorf("guest metrics: %w", err)
	}
	if !slices.ContainsFunc(inventory.Guests, func(g supervision.Guest) bool { return g.VMID == vmid }) {
		return domain.TimeSeries{}, domain.NewNotFound("guest", vmid)
	}
	if len(names) == 0 {
		names = supervision.AllGuestSeries()
	}
	series, err := m.series.GuestSeries(ctx, vmid, slices.Compact(slices.Sorted(slices.Values(names))), w)
	if err != nil {
		return domain.TimeSeries{}, fmt.Errorf("guest %d metrics: %w", vmid, err)
	}
	return domain.TimeSeries{Window: w, Series: series}, nil
}
