package availability

import (
	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
)

type AlertFilter struct {
	State    availability.AlertState
	Severity domain.Severity
	Guest    string
}

func (f AlertFilter) matches(a availability.Alert) bool {
	return (f.State == "" || a.State == f.State) &&
		(f.Severity == "" || a.Severity == f.Severity) &&
		(f.Guest == "" || a.Labels["vmid"] == f.Guest)
}
