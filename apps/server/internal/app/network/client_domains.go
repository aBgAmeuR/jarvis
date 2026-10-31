package network

import (
	"context"
	"fmt"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

func (c *Clients) Domains(ctx context.Context, id string, status network.DomainStatus, limit int, start, end time.Time) ([]network.DomainCount, error) {
	limit, err := domainLimit(limit)
	if err != nil {
		return nil, err
	}
	w, err := domain.NewWindow(start, end, c.now())
	if err != nil {
		return nil, err
	}
	client, _, err := c.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	domains, err := c.domains.ClientDomains(ctx, client.IP, status, w, limit)
	if err != nil {
		return nil, fmt.Errorf("client %s domains: %w", id, err)
	}
	return domains, nil
}
