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
	q := query{name: string(name)}
	switch name {
	case supervision.NodeCPUUsageRatio:
		q.unit, q.expr = domain.UnitRatio, "max(pve_cpu_usage_ratio"+byID+")"
	case supervision.NodeMemoryUsedBytes:
		q.unit, q.expr = domain.UnitBytes, "max(pve_memory_usage_bytes"+byID+")"
	case supervision.NodeDiskUsedBytes:
		q.unit, q.expr = domain.UnitBytes, "max(pve_disk_usage_bytes"+byID+")"
	case supervision.NodeCPUTemperature:
		q.unit, q.expr = domain.UnitCelsius, "max(node_hwmon_temp_celsius{instance="+strconv.Quote(node)+`,chip=~"platform_coretemp.*"}`+")"
	case supervision.NodeLoad1:
		q.unit, q.expr = domain.UnitLoad, "max(node_load1"+byInstance+")"
	default:
		return query{}, false
	}
	return q, true
}
