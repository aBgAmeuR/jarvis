package availability

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"time"

	"jarvis/server/internal/domain/availability"
)

// CertificateExpiry adds the time left, negative once the certificate expired.
type CertificateExpiry struct {
	availability.Certificate
	Remaining time.Duration
}

func (c *Certificates) List(ctx context.Context) ([]CertificateExpiry, error) {
	certificates, err := c.source.Certificates(ctx)
	if err != nil {
		return nil, fmt.Errorf("list certificates: %w", err)
	}
	now := c.now()
	out := make([]CertificateExpiry, 0, len(certificates))
	for _, certificate := range certificates {
		out = append(out, CertificateExpiry{Certificate: certificate, Remaining: certificate.ExpiresAt.Sub(now)})
	}
	slices.SortFunc(out, func(x, y CertificateExpiry) int { return cmp.Compare(x.Remaining, y.Remaining) })
	return out, nil
}
