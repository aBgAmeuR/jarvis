package logs

import "time"

type Entry struct {
	Timestamp time.Time
	Line      string
	Level     string
	Service   string
	Instance  string
	// VMID is zero when the stream has no vmid label.
	VMID int
}
