package supervision

import "jarvis/server/internal/domain"

type GuestAlert struct {
	VMID     int
	Severity domain.Severity
}
