package availability

import (
	"cmp"
	"context"
	"fmt"
	"slices"

	"jarvis/server/internal/domain/availability"
)

func (s *Silences) List(ctx context.Context, state availability.SilenceState) ([]availability.Silence, error) {
	all, err := s.source.Silences(ctx)
	if err != nil {
		return nil, fmt.Errorf("list silences: %w", err)
	}
	silences := slices.DeleteFunc(slices.Clone(all), func(silence availability.Silence) bool {
		return state != "" && silence.State != state
	})
	slices.SortFunc(silences, func(x, y availability.Silence) int {
		return cmp.Or(y.EndsAt.Compare(x.EndsAt), cmp.Compare(x.ID, y.ID))
	})
	return silences, nil
}
