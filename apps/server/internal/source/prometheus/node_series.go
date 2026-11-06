package prometheus

import (
	"context"
	"strconv"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/supervision"
)

func (c *Client) NodeSeries(ctx context.Context, node string, names []supervision.NodeSeries, w domain.Window) ([]domain.Series, error) {
	queries := make([]query, 0, len(names))
	for _, n := range names {
		if q, ok := nodeQuery(n, node); ok {
			queries = append(queries, q)
		}
	}
	return c.runQueries(ctx, queries, w)
}

// nodeQuery reads the pve-exporter series by id, and the host's own Alloy by instance.
func nodeQuery(name supervision.NodeSeries, node string) (query, bool) {
	byID := "{id=" + strconv.Quote("node/"+node) + "}"
	byInstance := "{instance=" + strconv.Quote(node) + "}"
	n := string(name)
	switch name {
	case supervision.NodeCPUUsageRatio:
		return gauge(n, domain.UnitRatio, "pve_cpu_usage_ratio"+byID), true
	case supervision.NodeMemoryUsedBytes:
		return gauge(n, domain.UnitBytes, "pve_memory_usage_bytes"+byID), true
	case supervision.NodeDiskUsedBytes:
		return gauge(n, domain.UnitBytes, "pve_disk_usage_bytes"+byID), true
	case supervision.NodeCPUTemperature:
		return gauge(n, domain.UnitCelsius, "node_hwmon_temp_celsius{instance="+strconv.Quote(node)+`,chip=~"platform_coretemp.*"}`), true
	case supervision.NodeLoad1:
		return gauge(n, domain.UnitLoad, "node_load1"+byInstance), true
	default:
		return query{}, false
	}
}
