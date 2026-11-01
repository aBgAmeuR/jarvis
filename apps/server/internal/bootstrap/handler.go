package bootstrap

import (
	"log/slog"
	"net/http"

	"jarvis/server/internal/httpapi"
	"jarvis/server/internal/httpapi/middleware"
	"jarvis/server/internal/httpapi/ops"
	"jarvis/server/internal/platform/telemetry"
	"jarvis/server/internal/platform/webui"
)

func newHandler(logger *slog.Logger, metrics *telemetry.Metrics, deps httpapi.Deps) http.Handler {
	mux := http.NewServeMux()
	ops.Register(mux, deps.Health, metrics.Handler())
	httpapi.NewAPI(mux, deps)
	mux.Handle("GET /", webui.Handler())
	return middleware.Chain(mux,
		telemetry.TraceServer,
		middleware.RequestID,
		middleware.Observe(middleware.AccessLog(logger), middleware.Metrics(metrics)),
		middleware.Recover(logger),
		middleware.SecurityHeaders,
		middleware.ETag,
	)
}
