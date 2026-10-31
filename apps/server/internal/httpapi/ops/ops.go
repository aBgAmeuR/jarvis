package ops

import (
	"net/http"

	apphealth "jarvis/server/internal/app/health"
)

// Register adds the operations paths, outside the versioned API and identity.
func Register(mux *http.ServeMux, health *apphealth.Service, metrics http.Handler) {
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		if !health.Ready(r.Context()) {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("GET /metrics", metrics)
}
