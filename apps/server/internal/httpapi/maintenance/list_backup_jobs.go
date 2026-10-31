package maintenance

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"

	appmaintenance "jarvis/server/internal/app/maintenance"
	"jarvis/server/internal/httpapi/problem"
)

type backupJobBody struct {
	ID        string     `json:"id" doc:"Proxmox backup job id." example:"backup-3f1a2b4c-0d5e"`
	Enabled   bool       `json:"enabled" doc:"Whether the job runs." example:"true"`
	Schedule  string     `json:"schedule" doc:"When the job runs, as Proxmox writes it." example:"03:00"`
	Storage   string     `json:"storage" doc:"Where the backups go." example:"local"`
	Mode      string     `json:"mode" doc:"Backup mode." example:"snapshot"`
	Comment   string     `json:"comment,omitempty" doc:"Comment on the job." example:"Nightly"`
	NextRunAt *time.Time `json:"nextRunAt,omitempty" doc:"Next run, RFC 3339 in UTC." example:"2026-11-01T03:00:00Z"`
}

type listBackupJobsOutput struct {
	Body struct {
		Items []backupJobBody `json:"items" nullable:"false" doc:"Backup jobs by id."`
	}
}

func registerListBackupJobs(api huma.API, problems *problem.Mapper, backups *appmaintenance.Backups) {
	huma.Register(api, huma.Operation{
		OperationID: "listBackupJobs",
		Method:      http.MethodGet,
		Path:        "/backup-jobs",
		Summary:     "List backup jobs",
		Tags:        []string{"maintenance"},
	}, func(ctx context.Context, _ *struct{}) (*listBackupJobsOutput, error) {
		jobs, err := backups.Jobs(ctx)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listBackupJobsOutput{}
		out.Body.Items = make([]backupJobBody, 0, len(jobs))
		for _, j := range jobs {
			out.Body.Items = append(out.Body.Items, backupJobBody{
				ID: j.ID, Enabled: j.Enabled, Schedule: j.Schedule, Storage: j.Storage,
				Mode: j.Mode, Comment: j.Comment, NextRunAt: j.NextRunAt,
			})
		}
		return out, nil
	})
}
