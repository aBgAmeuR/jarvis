package proxmox

import (
	"context"
	"slices"
	"time"

	"jarvis/server/internal/domain/maintenance"
	"jarvis/server/internal/source/failure"
)

func (c *Client) BackupVolumes(ctx context.Context) ([]maintenance.Volume, error) {
	inventory, err := c.Inventory(ctx)
	if err != nil {
		return nil, err
	}
	var volumes []maintenance.Volume
	for _, s := range inventory.Storages {
		if s.Status != "available" || !slices.Contains(s.Content, "backup") {
			continue
		}
		var body struct {
			Data []struct {
				VMID  int    `json:"vmid"`
				Size  uint64 `json:"size"`
				CTime int64  `json:"ctime"`
			} `json:"data"`
		}
		if err := c.fetcher.GetJSON(ctx, "/api2/json/nodes/"+s.Node+"/storage/"+s.Name+"/content?content=backup", &body); err != nil {
			return nil, failure.Wrap(name, err)
		}
		for _, v := range body.Data {
			volumes = append(volumes, maintenance.Volume{VMID: v.VMID, SizeBytes: v.Size, CreatedAt: time.Unix(v.CTime, 0).UTC()})
		}
	}
	return volumes, nil
}
