package maintenance

import (
	"time"

	domainmaintenance "jarvis/server/internal/domain/maintenance"
)

type backupBody struct {
	VMID                int        `json:"vmid" doc:"Proxmox guest id." example:"116"`
	Name                string     `json:"name" doc:"Guest name." example:"dns"`
	Covered             bool       `json:"covered" doc:"Whether a backup job covers the guest." example:"true"`
	BackupCount         int        `json:"backupCount" doc:"Backup files kept for the guest." example:"7"`
	LastBackupAt        *time.Time `json:"lastBackupAt,omitempty" doc:"When the newest backup file was written, RFC 3339 in UTC. Absent without any." example:"2026-10-31T03:00:00Z"`
	LastBackupSizeBytes uint64     `json:"lastBackupSizeBytes" doc:"Size of the newest backup file." example:"524288000"`
	LastRunAt           *time.Time `json:"lastRunAt,omitempty" doc:"When the last backup run started, RFC 3339 in UTC. Absent without any." example:"2026-10-31T03:00:00Z"`
	LastResult          string     `json:"lastResult" doc:"Result of the last backup run." enum:"ok,warning,error,unknown" example:"ok"`
}

func newBackupBody(b domainmaintenance.Backup) backupBody {
	return backupBody{
		VMID:                b.Guest.VMID,
		Name:                b.Guest.Name,
		Covered:             b.Covered,
		BackupCount:         b.Count,
		LastBackupAt:        b.LastBackupAt,
		LastBackupSizeBytes: b.LastSizeBytes,
		LastRunAt:           b.LastRunAt,
		LastResult:          string(b.LastResult),
	}
}
