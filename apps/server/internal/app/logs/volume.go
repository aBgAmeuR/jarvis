package logs

import (
	"context"
	"fmt"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/logs"
)

// Volume ignores the limit of q.
func (l *Logs) Volume(ctx context.Context, q logs.Query) (logs.Volume, error) {
	w, err := domain.NewWindow(q.Start, q.End, l.now())
	if err != nil {
		return logs.Volume{}, err
	}
	counts, err := l.source.LogCounts(ctx, q, w)
	if err != nil {
		return logs.Volume{}, fmt.Errorf("log volume: %w", err)
	}
	return logs.Volume{Window: w, Counts: counts}, nil
}
