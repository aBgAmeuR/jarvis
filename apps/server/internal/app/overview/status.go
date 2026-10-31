package overview

import (
	"cmp"
	"slices"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/overview"
)

func statusOf(alerts []overview.Alert) overview.Status {
	if len(alerts) == 0 {
		return overview.StatusOK
	}
	worst := slices.MinFunc(alerts, func(a, b overview.Alert) int { return cmp.Compare(a.Severity.Rank(), b.Severity.Rank()) })
	switch worst.Severity {
	case domain.SeverityCritical:
		return overview.StatusCritical
	case domain.SeverityWarning:
		return overview.StatusWarning
	default:
		return overview.StatusInfo
	}
}
