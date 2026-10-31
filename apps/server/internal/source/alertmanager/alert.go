package alertmanager

import (
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
)

type alert struct {
	Fingerprint string            `json:"fingerprint"`
	StartsAt    time.Time         `json:"startsAt"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	Status      struct {
		State string `json:"state"`
	} `json:"status"`
}

func (a alert) toDomain() availability.Alert {
	state := availability.AlertFiring
	if a.Status.State == "suppressed" {
		state = availability.AlertSuppressed
	}
	return availability.Alert{
		Fingerprint: a.Fingerprint,
		Name:        a.Labels["alertname"],
		Severity:    domain.ParseSeverity(a.Labels["severity"]),
		State:       state,
		StartsAt:    a.StartsAt,
		Labels:      a.Labels,
		Annotations: a.Annotations,
	}
}
