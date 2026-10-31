package logs

import (
	"context"

	"jarvis/server/internal/domain/logs"
)

// Logs returns at most Limit lines, newest first.
type LogSource interface {
	Logs(ctx context.Context, q logs.Query) ([]logs.Entry, error)
}
