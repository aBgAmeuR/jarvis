package proxmox

import (
	"context"

	"jarvis/server/internal/domain/supervision"
	"jarvis/server/internal/source/failure"
)

func (c *Client) Inventory(ctx context.Context) (supervision.Inventory, error) {
	var body struct {
		Data []resource `json:"data"`
	}
	if err := c.fetcher.GetJSON(ctx, "/api2/json/cluster/resources", &body); err != nil {
		return supervision.Inventory{}, failure.Wrap(name, err)
	}

	var inventory supervision.Inventory
	for _, r := range body.Data {
		switch r.Type {
		case "node":
			inventory.Nodes = append(inventory.Nodes, r.node())
		case "storage":
			inventory.Storages = append(inventory.Storages, r.storage())
		case "lxc", "qemu":
			if r.Template != 1 {
				inventory.Guests = append(inventory.Guests, r.guest())
			}
		}
	}
	return inventory, nil
}
