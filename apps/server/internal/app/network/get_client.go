package network

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

func (c *Clients) Get(ctx context.Context, id string) (network.Client, []domain.Degraded, error) {
	list, err := c.List(ctx, "", 0)
	if err != nil {
		return network.Client{}, nil, err
	}
	for _, client := range list.Items {
		if client.ID == id {
			return client, list.Degraded, nil
		}
	}
	return network.Client{}, nil, domain.NewNotFound("client", id)
}
