package supervision

import "time"

type Metrics struct {
	inventory InventorySource
	series    MetricsSource
	now       func() time.Time
}

func NewMetrics(inventory InventorySource, series MetricsSource) *Metrics {
	return &Metrics{inventory: inventory, series: series, now: time.Now}
}
