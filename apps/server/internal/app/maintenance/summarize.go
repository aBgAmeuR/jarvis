package maintenance

import "jarvis/server/internal/domain/maintenance"

func summarize(guest maintenance.GuestRef, covered bool, volumes []maintenance.Volume, tasks []maintenance.Task) maintenance.Backup {
	backup := maintenance.Backup{Guest: guest, Covered: covered, LastResult: maintenance.BackupUnknown}
	for _, v := range volumes {
		if v.VMID != guest.VMID {
			continue
		}
		backup.Count++
		if backup.LastBackupAt == nil || v.CreatedAt.After(*backup.LastBackupAt) {
			at := v.CreatedAt
			backup.LastBackupAt, backup.LastSizeBytes = &at, v.SizeBytes
		}
	}
	for _, t := range tasks {
		if t.VMID == guest.VMID && (backup.LastRunAt == nil || t.StartedAt.After(*backup.LastRunAt)) {
			at := t.StartedAt
			backup.LastRunAt, backup.LastResult = &at, t.Result
		}
	}
	return backup
}
