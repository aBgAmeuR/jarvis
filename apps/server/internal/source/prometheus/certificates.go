package prometheus

import (
	"context"
	"slices"
	"strconv"
	"strings"
	"time"

	"jarvis/server/internal/domain/availability"
	"jarvis/server/internal/source/failure"
)

func (c *Client) Certificates(ctx context.Context) ([]availability.Certificate, error) {
	entries, err := c.queryVector(ctx, `probe_ssl_earliest_cert_expiry{job="probe_http"}`)
	if err != nil {
		return nil, failure.Wrap(name, err)
	}
	certificates := make([]availability.Certificate, 0, len(entries))
	for _, e := range entries {
		certificate := availability.Certificate{
			Site:      e.labels["site"],
			Target:    e.labels["target"],
			ExpiresAt: time.Unix(int64(e.value), 0).UTC(),
		}
		if vmid, err := strconv.Atoi(e.labels["vmid"]); err == nil {
			certificate.VMID = &vmid
		}
		certificates = append(certificates, certificate)
	}
	slices.SortFunc(certificates, func(a, b availability.Certificate) int { return strings.Compare(a.Site, b.Site) })
	return certificates, nil
}
