package availability

import (
	"time"

	domainavailability "jarvis/server/internal/domain/availability"
)

type matcherBody struct {
	Name    string `json:"name" doc:"Label name." example:"alertname"`
	Value   string `json:"value" doc:"Label value or pattern." example:"GuestNotBackedUp"`
	IsRegex bool   `json:"isRegex" doc:"Whether the value is a regular expression." example:"false"`
	IsEqual bool   `json:"isEqual" doc:"Whether the matcher selects equal labels, not different ones." example:"true"`
}

type silenceBody struct {
	ID        string        `json:"id" doc:"Alertmanager silence id." example:"f3b1c0aa-3f3d-4a0e-9a9e-7c9d5d3a1b11"`
	State     string        `json:"state" doc:"Silence state." enum:"active,pending,expired" example:"active"`
	StartsAt  time.Time     `json:"startsAt" doc:"When the silence starts, RFC 3339 in UTC." example:"2026-10-31T10:00:00Z"`
	EndsAt    time.Time     `json:"endsAt" doc:"When the silence ends, RFC 3339 in UTC." example:"2026-11-01T10:00:00Z"`
	CreatedBy string        `json:"createdBy" doc:"Who created the silence." example:"antoine"`
	Comment   string        `json:"comment" doc:"Why it was created." example:"Planned backup rollout"`
	Matchers  []matcherBody `json:"matchers" nullable:"false" doc:"Labels the silence applies to."`
}

func newSilenceBody(s domainavailability.Silence) silenceBody {
	body := silenceBody{
		ID:        s.ID,
		State:     string(s.State),
		StartsAt:  s.StartsAt,
		EndsAt:    s.EndsAt,
		CreatedBy: s.CreatedBy,
		Comment:   s.Comment,
		Matchers:  make([]matcherBody, 0, len(s.Matchers)),
	}
	for _, m := range s.Matchers {
		body.Matchers = append(body.Matchers, matcherBody{Name: m.Name, Value: m.Value, IsRegex: m.IsRegex, IsEqual: m.IsEqual})
	}
	return body
}
