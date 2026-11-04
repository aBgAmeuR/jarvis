package domain

// Grafana is where Explore links point. URL is empty when the links are off.
type Grafana struct {
	URL           string
	PrometheusUID string
	LokiUID       string
}
