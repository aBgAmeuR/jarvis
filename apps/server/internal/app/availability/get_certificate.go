package availability

import (
	"context"

	"jarvis/server/internal/domain"
)

func (c *Certificates) Get(ctx context.Context, site string) (CertificateExpiry, error) {
	certificates, err := c.List(ctx)
	if err != nil {
		return CertificateExpiry{}, err
	}
	for _, certificate := range certificates {
		if certificate.Site == site {
			return certificate, nil
		}
	}
	return CertificateExpiry{}, domain.NewNotFound("certificate", site)
}
