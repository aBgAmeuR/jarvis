package telemetry

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	registry *prometheus.Registry

	requests       *prometheus.CounterVec
	requestSeconds *prometheus.HistogramVec
	sourceRequests *prometheus.CounterVec
	sourceSeconds  *prometheus.HistogramVec
}

func NewMetrics(version, commit string) *Metrics {
	m := &Metrics{
		registry: prometheus.NewRegistry(),
		requests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "jarvis_http_requests_total",
			Help: "Requests served, by method, route and status code.",
		}, []string{"method", "route", "code"}),
		requestSeconds: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "jarvis_http_request_duration_seconds",
			Help:    "Time to serve a request, by method and route.",
			Buckets: prometheus.DefBuckets,
		}, []string{"method", "route"}),
		sourceRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "jarvis_source_requests_total",
			Help: "Requests sent to data sources, by host and status code (error when none).",
		}, []string{"host", "code"}),
		sourceSeconds: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "jarvis_source_request_duration_seconds",
			Help:    "Time to get an answer from a data source, by host.",
			Buckets: prometheus.DefBuckets,
		}, []string{"host"}),
	}
	buildInfo := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "jarvis_build_info",
		Help:        "Build information, always 1.",
		ConstLabels: prometheus.Labels{"version": version, "commit": commit},
	})
	buildInfo.Set(1)
	m.registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		buildInfo, m.requests, m.requestSeconds, m.sourceRequests, m.sourceSeconds,
	)
	return m
}

func (m *Metrics) Handler() http.Handler {
	return promhttp.HandlerFor(m.registry, promhttp.HandlerOpts{Registry: m.registry})
}
