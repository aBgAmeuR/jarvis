package loki

import (
	"context"
	"strconv"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

func (c *Client) ClientDomains(ctx context.Context, ip string, status network.DomainStatus, w domain.Window, limit int) ([]network.DomainCount, error) {
	return c.topDomains(ctx, `{service_name="blocky"} | client_ip=`+strconv.Quote(ip), status, w, limit)
}
