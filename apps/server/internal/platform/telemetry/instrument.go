package telemetry

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func TraceServer(next http.Handler) http.Handler {
	return otelhttp.NewHandler(next, "http.server", otelhttp.WithSpanNameFormatter(
		func(_ string, r *http.Request) string { return r.Method + " " + r.URL.Path },
	))
}

func TraceClient(next http.RoundTripper) http.RoundTripper {
	return otelhttp.NewTransport(next)
}
