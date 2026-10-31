package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
)

// The query string and the identity headers are never logged.
func AccessLog(logger *slog.Logger) Observer {
	return func(ctx context.Context, o Observation) {
		level := slog.LevelInfo
		switch {
		case o.Status >= http.StatusInternalServerError:
			level = slog.LevelError
		case !strings.HasPrefix(o.Path, "/api/"):
			level = slog.LevelDebug
		}
		logger.Log(ctx, level, "request",
			"method", o.Method,
			"path", o.Path,
			"route", o.Route,
			"status", o.Status,
			"bytes", o.Bytes,
			"duration_ms", float64(o.Duration.Microseconds())/1000,
		)
	}
}
