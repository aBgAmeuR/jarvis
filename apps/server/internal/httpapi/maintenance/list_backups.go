package maintenance

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appmaintenance "jarvis/server/internal/app/maintenance"
	"jarvis/server/internal/httpapi/problem"
)

type listBackupsInput struct {
	Guest int `query:"guest" doc:"Only this guest VMID." example:"116"`
}

type listBackupsOutput struct {
	Body struct {
		Items []backupBody `json:"items" nullable:"false" doc:"Backup status of the guests by VMID."`
	}
}

func registerListBackups(api huma.API, problems *problem.Mapper, backups *appmaintenance.Backups) {
	huma.Register(api, huma.Operation{
		OperationID: "listBackups",
		Method:      http.MethodGet,
		Path:        "/backups",
		Summary:     "List the backup status of guests",
		Tags:        []string{"maintenance"},
	}, func(ctx context.Context, in *listBackupsInput) (*listBackupsOutput, error) {
		list, err := backups.List(ctx, in.Guest)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listBackupsOutput{}
		out.Body.Items = make([]backupBody, 0, len(list))
		for _, b := range list {
			out.Body.Items = append(out.Body.Items, newBackupBody(b))
		}
		return out, nil
	})
}
