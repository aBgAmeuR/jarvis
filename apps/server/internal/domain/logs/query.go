package logs

import "time"

type Query struct {
	VMID     int
	Service  string
	Level    string
	Contains string
	// Start is inclusive and End exclusive.
	Start time.Time
	End   time.Time
	Limit int
}
