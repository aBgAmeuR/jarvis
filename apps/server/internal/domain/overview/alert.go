package overview

import (
	"time"

	"jarvis/server/internal/domain"
)

type Alert struct {
	Fingerprint string
	Name        string
	Severity    domain.Severity
	StartsAt    time.Time
	Summary     string
}
