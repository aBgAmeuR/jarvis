package availability

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
)

func (s *Silences) Get(ctx context.Context, id string) (availability.Silence, error) {
	silences, err := s.List(ctx, "")
	if err != nil {
		return availability.Silence{}, err
	}
	for _, silence := range silences {
		if silence.ID == id {
			return silence, nil
		}
	}
	return availability.Silence{}, domain.NewNotFound("silence", id)
}
