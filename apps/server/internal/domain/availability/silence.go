package availability

import "time"

type SilenceState string

const (
	SilenceActive  SilenceState = "active"
	SilencePending SilenceState = "pending"
	SilenceExpired SilenceState = "expired"
)

type Matcher struct {
	Name    string
	Value   string
	IsRegex bool
	IsEqual bool
}

type Silence struct {
	ID        string
	State     SilenceState
	StartsAt  time.Time
	EndsAt    time.Time
	CreatedBy string
	Comment   string
	Matchers  []Matcher
}
