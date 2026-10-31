package telemetry

import (
	"strconv"
	"time"
)

// route is the matched pattern, "" when none.
func (m *Metrics) ObserveRequest(method, route string, status int, elapsed time.Duration) {
	if route == "" {
		route = "unmatched"
	}
	m.requests.WithLabelValues(method, route, strconv.Itoa(status)).Inc()
	m.requestSeconds.WithLabelValues(method, route).Observe(elapsed.Seconds())
}
