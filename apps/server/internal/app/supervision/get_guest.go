package supervision

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/supervision"
)

// Get reads the whole list, which the HTTP client cache makes cheap.
func (g *Guests) Get(ctx context.Context, vmid int) (supervision.Guest, []domain.Degraded, error) {
	list, err := g.List(ctx)
	if err != nil {
		return supervision.Guest{}, nil, err
	}
	for _, guest := range list.Items {
		if guest.VMID == vmid {
			return guest, list.Degraded, nil
		}
	}
	return supervision.Guest{}, nil, domain.NewNotFound("guest", vmid)
}
