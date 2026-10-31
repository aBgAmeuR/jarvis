package bootstrap

import (
	appmaintenance "jarvis/server/internal/app/maintenance"
	httpmaintenance "jarvis/server/internal/httpapi/maintenance"
)

func newMaintenance(s sources) httpmaintenance.Services {
	return httpmaintenance.Services{
		Backups: appmaintenance.NewBackups(s.proxmox, s.proxmox),
		Drift:   appmaintenance.NewDrift(s.proxmox, s.catalog, s.proxmox),
	}
}
