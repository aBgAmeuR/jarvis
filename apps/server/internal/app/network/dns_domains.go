package network

import (
	"context"
	"fmt"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

func (d *DNS) Domains(ctx context.Context, status network.DomainStatus, limit int, start, end time.Time) ([]network.DomainCount, error) {
	limit, err := domainLimit(limit)
	if err != nil {
		return nil, err
	}
	w, err := domain.NewWindow(start, end, d.now())
	if err != nil {
		return nil, err
	}
	domains, err := d.domains.DNSDomains(ctx, status, w, limit)
	if err != nil {
		return nil, fmt.Errorf("dns domains: %w", err)
	}
	return domains, nil
}
