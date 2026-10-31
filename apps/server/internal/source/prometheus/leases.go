package prometheus

import (
	"context"
	"slices"
	"strings"
	"time"

	"jarvis/server/internal/domain/network"
	"jarvis/server/internal/source/failure"
)

func (c *Client) Leases(ctx context.Context) ([]network.Lease, error) {
	entries, err := c.queryVector(ctx, "dnsmasq_lease_expiry_timestamp_seconds")
	if err != nil {
		return nil, failure.Wrap(name, err)
	}
	leases := make([]network.Lease, 0, len(entries))
	for _, e := range entries {
		leases = append(leases, network.Lease{
			MAC:       strings.ToLower(e.labels["mac"]),
			IP:        e.labels["ip"],
			ExpiresAt: time.Unix(int64(e.value), 0).UTC(),
		})
	}
	slices.SortFunc(leases, func(a, b network.Lease) int { return strings.Compare(a.MAC, b.MAC) })
	return leases, nil
}
