package prometheus

import (
	"context"
	"fmt"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/supervision"
)

func (c *Client) GuestSeries(ctx context.Context, vmid int, names []supervision.GuestSeries, w domain.Window) ([]domain.Series, error) {
	queries := make([]query, 0, len(names))
	for _, n := range names {
		if q, ok := guestQuery(n, vmid, w); ok {
			queries = append(queries, q)
		}
	}
	return c.runQueries(ctx, queries, w)
}

// guestQuery uses the pve-exporter series, which carry the vmid label.
func guestQuery(name supervision.GuestSeries, vmid int, w domain.Window) (query, bool) {
	selector := fmt.Sprintf(`{vmid="%d"}`, vmid)
	rate := rateRange(w)
	q := query{name: string(name)}
	switch name {
	case supervision.GuestCPUUsageRatio:
		q = gauge(q.name, domain.UnitRatio, "pve_cpu_usage_ratio"+selector)
	case supervision.GuestMemoryUsedBytes:
		q = gauge(q.name, domain.UnitBytes, "pve_memory_usage_bytes"+selector)
	case supervision.GuestDiskUsedBytes:
		q = gauge(q.name, domain.UnitBytes, "pve_disk_usage_bytes"+selector)
	case supervision.GuestDiskReadBytesPerSec:
		q.unit, q.expr = domain.UnitBytesPerSecond, "max(rate(pve_disk_read_bytes_total"+selector+rate+"))"
	case supervision.GuestDiskWriteBytesPerSec:
		q.unit, q.expr = domain.UnitBytesPerSecond, "max(rate(pve_disk_written_bytes_total"+selector+rate+"))"
	case supervision.GuestNetworkReceivePerSec:
		q.unit, q.expr = domain.UnitBytesPerSecond, "max(rate(pve_network_receive_bytes_total"+selector+rate+"))"
	case supervision.GuestNetworkTransmitPerSec:
		q.unit, q.expr = domain.UnitBytesPerSecond, "max(rate(pve_network_transmit_bytes_total"+selector+rate+"))"
	default:
		return query{}, false
	}
	return q, true
}
