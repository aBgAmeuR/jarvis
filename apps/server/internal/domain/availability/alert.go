package availability

import (
	"time"

	"jarvis/server/internal/domain"
)

type AlertState string

const (
	AlertFiring     AlertState = "firing"
	AlertSuppressed AlertState = "suppressed"
	AlertResolved   AlertState = "resolved"
)

type Alert struct {
	Fingerprint string
	Name        string
	Severity    domain.Severity
	State       AlertState
	StartsAt    time.Time
	EndsAt      *time.Time
	Labels      map[string]string
	Annotations map[string]string
}
