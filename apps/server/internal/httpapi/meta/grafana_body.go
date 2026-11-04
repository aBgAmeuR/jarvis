package meta

type GrafanaBody struct {
	URL           string `json:"url" doc:"Grafana base URL, for Explore links." example:"https://monitoring.antoinejosset.fr"`
	PrometheusUID string `json:"prometheusUid" doc:"Uid of the Prometheus data source in Grafana." example:"prometheus"`
	LokiUID       string `json:"lokiUid" doc:"Uid of the Loki data source in Grafana." example:"loki"`
}
