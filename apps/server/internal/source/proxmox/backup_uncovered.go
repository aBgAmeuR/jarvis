package proxmox

import (
	"context"

	"jarvis/server/internal/source/failure"
)

func (c *Client) UncoveredGuests(ctx context.Context) ([]int, error) {
	var body struct {
		Data []struct {
			VMID int `json:"vmid"`
		} `json:"data"`
	}
	if err := c.fetcher.GetJSON(ctx, "/api2/json/cluster/backup-info/not-backed-up", &body); err != nil {
		return nil, failure.Wrap(name, err)
	}
	vmids := make([]int, 0, len(body.Data))
	for _, g := range body.Data {
		vmids = append(vmids, g.VMID)
	}
	return vmids, nil
}
