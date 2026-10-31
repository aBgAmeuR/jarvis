package alertmanager

import (
	"time"

	"jarvis/server/internal/domain/availability"
)

type silence struct {
	ID        string    `json:"id"`
	StartsAt  time.Time `json:"startsAt"`
	EndsAt    time.Time `json:"endsAt"`
	CreatedBy string    `json:"createdBy"`
	Comment   string    `json:"comment"`
	Status    struct {
		State string `json:"state"`
	} `json:"status"`
	Matchers []struct {
		Name    string `json:"name"`
		Value   string `json:"value"`
		IsRegex bool   `json:"isRegex"`
		IsEqual *bool  `json:"isEqual"`
	} `json:"matchers"`
}

func (s silence) toDomain() availability.Silence {
	out := availability.Silence{
		ID:        s.ID,
		State:     availability.SilenceState(s.Status.State),
		StartsAt:  s.StartsAt,
		EndsAt:    s.EndsAt,
		CreatedBy: s.CreatedBy,
		Comment:   s.Comment,
		Matchers:  make([]availability.Matcher, 0, len(s.Matchers)),
	}
	for _, m := range s.Matchers {
		// Alertmanager omits isEqual on equality matchers.
		equal := m.IsEqual == nil || *m.IsEqual
		out.Matchers = append(out.Matchers, availability.Matcher{Name: m.Name, Value: m.Value, IsRegex: m.IsRegex, IsEqual: equal})
	}
	return out
}
