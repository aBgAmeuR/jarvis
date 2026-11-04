package logs

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/logs"
)

// Logs returns at most Limit lines, newest first.
type LogSource interface {
	Logs(ctx context.Context, q logs.Query) ([]logs.Entry, error)
	// LogCounts counts lines per level in each step of the window.
	LogCounts(ctx context.Context, q logs.Query, w domain.Window) ([]logs.Count, error)
}
