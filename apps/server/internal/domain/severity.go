package domain

type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityWarning  Severity = "warning"
	SeverityInfo     Severity = "info"
	SeverityNone     Severity = "none"
)

// ParseSeverity maps anything else, such as the "none" of Watchdog, to SeverityNone.
func ParseSeverity(s string) Severity {
	switch v := Severity(s); v {
	case SeverityCritical, SeverityWarning, SeverityInfo:
		return v
	default:
		return SeverityNone
	}
}

// Rank is 0 for the most severe.
func (s Severity) Rank() int {
	switch s {
	case SeverityCritical:
		return 0
	case SeverityWarning:
		return 1
	case SeverityInfo:
		return 2
	default:
		return 3
	}
}
