package proxmox

import (
	"context"

	"jarvis/server/internal/domain/maintenance"
)

func (c *Client) BackupGuests(ctx context.Context) ([]maintenance.GuestRef, error) {
	inventory, err := c.Inventory(ctx)
	if err != nil {
		return nil, err
	}
	guests := make([]maintenance.GuestRef, 0, len(inventory.Guests))
	for _, g := range inventory.Guests {
		guests = append(guests, maintenance.GuestRef{VMID: g.VMID, Name: g.Name, Kind: string(g.Kind)})
	}
	return guests, nil
}
