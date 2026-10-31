package loki

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

func (c *Client) DNSDomains(ctx context.Context, status network.DomainStatus, w domain.Window, limit int) ([]network.DomainCount, error) {
	return c.topDomains(ctx, `{service_name="blocky"}`, status, w, limit)
}
