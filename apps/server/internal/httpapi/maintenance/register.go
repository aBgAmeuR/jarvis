package maintenance

import (
	"github.com/danielgtaylor/huma/v2"

	appmaintenance "jarvis/server/internal/app/maintenance"
	"jarvis/server/internal/httpapi/problem"
)

type Services struct {
	Backups *appmaintenance.Backups
	Drift   *appmaintenance.Drift
}

func Register(api huma.API, problems *problem.Mapper, s Services) {
	registerListBackups(api, problems, s.Backups)
	registerListBackupJobs(api, problems, s.Backups)
	registerListDriftItems(api, problems, s.Drift)
}
