package telemetry

import (
	"net/http"
	"strconv"
	"time"
)

type sourceTransport struct {
	next    http.RoundTripper
	metrics *Metrics
}

func (m *Metrics) MeasureSources(next http.RoundTripper) http.RoundTripper {
	return sourceTransport{next: next, metrics: m}
}

func (t sourceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := t.next.RoundTrip(req)
	code := "error"
	if err == nil {
		code = strconv.Itoa(resp.StatusCode)
	}
	t.metrics.sourceRequests.WithLabelValues(req.URL.Host, code).Inc()
	t.metrics.sourceSeconds.WithLabelValues(req.URL.Host).Observe(time.Since(start).Seconds())
	return resp, err
}
