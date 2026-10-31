package availability

import (
	"time"

	domainavailability "jarvis/server/internal/domain/availability"
)

type alertBody struct {
	Fingerprint string            `json:"fingerprint" doc:"Alertmanager fingerprint, the alert id." example:"3a1f9c0d5b7e2a44"`
	Name        string            `json:"name" doc:"Alert rule name." example:"GuestDown"`
	Severity    string            `json:"severity" doc:"Severity label of the rule." enum:"critical,warning,info,none" example:"critical"`
	State       string            `json:"state" doc:"firing, or suppressed when silenced or inhibited." enum:"firing,suppressed,resolved" example:"firing"`
	StartsAt    time.Time         `json:"startsAt" doc:"When the alert started firing, RFC 3339 in UTC. For a resolved alert that began before the window, its first sample in the window." example:"2026-10-31T09:41:00Z"`
	EndsAt      *time.Time        `json:"endsAt,omitempty" doc:"When a resolved alert stopped firing, RFC 3339 in UTC." example:"2026-10-31T10:12:00Z"`
	Labels      map[string]string `json:"labels" doc:"Alert labels." example:"{\"vmid\":\"116\",\"severity\":\"critical\"}"`
	Annotations map[string]string `json:"annotations" doc:"Alert annotations, with summary and description." example:"{\"summary\":\"Guest 116 (dns) is not running\"}"`
}

func newAlertBody(a domainavailability.Alert) alertBody {
	return alertBody{
		Fingerprint: a.Fingerprint,
		Name:        a.Name,
		Severity:    string(a.Severity),
		State:       string(a.State),
		StartsAt:    a.StartsAt,
		EndsAt:      a.EndsAt,
		Labels:      a.Labels,
		Annotations: a.Annotations,
	}
}
