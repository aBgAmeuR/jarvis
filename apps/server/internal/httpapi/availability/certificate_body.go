package availability

import (
	"time"

	appavailability "jarvis/server/internal/app/availability"
)

type certificateBody struct {
	Site             string    `json:"site" doc:"Site host name." example:"files.antoinejosset.fr"`
	Target           string    `json:"target" doc:"URL probed." example:"https://files.antoinejosset.fr/healthz"`
	VMID             *int      `json:"vmid,omitempty" doc:"Guest behind the site." example:"115"`
	ExpiresAt        time.Time `json:"expiresAt" doc:"When the certificate expires, RFC 3339 in UTC." example:"2027-01-15T10:00:00Z"`
	ExpiresInSeconds int64     `json:"expiresInSeconds" doc:"Time left before it expires. Negative once expired." example:"7776000"`
}

func newCertificateBody(c appavailability.CertificateExpiry) certificateBody {
	return certificateBody{
		Site:             c.Site,
		Target:           c.Target,
		VMID:             c.VMID,
		ExpiresAt:        c.ExpiresAt,
		ExpiresInSeconds: int64(c.Remaining.Seconds()),
	}
}
