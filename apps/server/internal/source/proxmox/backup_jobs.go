package proxmox

import (
	"context"
	"time"

	"jarvis/server/internal/domain/maintenance"
	"jarvis/server/internal/source/failure"
)

type backupJob struct {
	ID       string `json:"id"`
	Enabled  *flag  `json:"enabled"`
	Schedule string `json:"schedule"`
	Storage  string `json:"storage"`
	Mode     string `json:"mode"`
	Comment  string `json:"comment"`
	NextRun  int64  `json:"next-run"`
}

func (c *Client) BackupJobs(ctx context.Context) ([]maintenance.BackupJob, error) {
	var body struct {
		Data []backupJob `json:"data"`
	}
	if err := c.fetcher.GetJSON(ctx, "/api2/json/cluster/backup", &body); err != nil {
		return nil, failure.Wrap(name, err)
	}
	jobs := make([]maintenance.BackupJob, 0, len(body.Data))
	for _, j := range body.Data {
		job := maintenance.BackupJob{
			ID:       j.ID,
			Enabled:  j.Enabled == nil || bool(*j.Enabled),
			Schedule: j.Schedule,
			Storage:  j.Storage,
			Mode:     j.Mode,
			Comment:  j.Comment,
		}
		if j.NextRun > 0 {
			next := time.Unix(j.NextRun, 0).UTC()
			job.NextRunAt = &next
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
