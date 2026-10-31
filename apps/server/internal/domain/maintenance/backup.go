package maintenance

import "time"

type BackupJob struct {
	ID        string
	Enabled   bool
	Schedule  string
	Storage   string
	Mode      string
	Comment   string
	NextRunAt *time.Time
}

type Volume struct {
	VMID      int
	CreatedAt time.Time
	SizeBytes uint64
}

type BackupResult string

const (
	BackupOK      BackupResult = "ok"
	BackupWarning BackupResult = "warning"
	BackupError   BackupResult = "error"
	BackupUnknown BackupResult = "unknown"
)

type Task struct {
	VMID      int
	StartedAt time.Time
	Result    BackupResult
}

type Backup struct {
	Guest         GuestRef
	Covered       bool
	Count         int
	LastBackupAt  *time.Time
	LastSizeBytes uint64
	LastRunAt     *time.Time
	LastResult    BackupResult
}
