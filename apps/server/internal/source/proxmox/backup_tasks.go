package proxmox

import (
	"context"
	"net/url"
	"strconv"
	"strings"
	"time"

	"jarvis/server/internal/domain/maintenance"
	"jarvis/server/internal/source/failure"
)

const taskLimit = 500

func (c *Client) BackupTasks(ctx context.Context) ([]maintenance.Task, error) {
	inventory, err := c.Inventory(ctx)
	if err != nil {
		return nil, err
	}
	var tasks []maintenance.Task
	for _, n := range inventory.Nodes {
		var body struct {
			Data []struct {
				ID        string `json:"id"`
				Status    string `json:"status"`
				StartTime int64  `json:"starttime"`
			} `json:"data"`
		}
		path := "/api2/json/nodes/" + url.PathEscape(n.Name) + "/tasks?typefilter=vzdump&limit=" + strconv.Itoa(taskLimit)
		if err := c.fetcher.GetJSON(ctx, path, &body); err != nil {
			return nil, failure.Wrap(name, err)
		}
		for _, t := range body.Data {
			vmid, err := strconv.Atoi(t.ID)
			if err != nil || t.Status == "" {
				continue
			}
			tasks = append(tasks, maintenance.Task{VMID: vmid, StartedAt: time.Unix(t.StartTime, 0).UTC(), Result: backupResult(t.Status)})
		}
	}
	return tasks, nil
}

// A finished task has the status OK, WARNINGS: n, or an error message.
func backupResult(status string) maintenance.BackupResult {
	switch {
	case status == "OK":
		return maintenance.BackupOK
	case strings.HasPrefix(status, "WARNINGS"):
		return maintenance.BackupWarning
	default:
		return maintenance.BackupError
	}
}
