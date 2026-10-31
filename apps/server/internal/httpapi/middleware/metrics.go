package middleware

import (
	"context"

	"jarvis/server/internal/platform/telemetry"
)

func Metrics(metrics *telemetry.Metrics) Observer {
	return func(_ context.Context, o Observation) {
		metrics.ObserveRequest(o.Method, o.Route, o.Status, o.Duration)
	}
}
