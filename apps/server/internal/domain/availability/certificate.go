package availability

import "time"

type Certificate struct {
	Site      string
	Target    string
	VMID      *int
	ExpiresAt time.Time
}
