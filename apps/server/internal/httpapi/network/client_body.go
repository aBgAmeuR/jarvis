package network

import (
	"time"

	domainnetwork "jarvis/server/internal/domain/network"
)

// ClientBody is exported so ClientDetailBody can embed it in the OpenAPI schema.
type ClientBody struct {
	ID             string     `json:"id" doc:"MAC address in lowercase, with hyphens." example:"bc-07-1d-35-46-1f"`
	MAC            string     `json:"mac" doc:"MAC address." example:"bc:07:1d:35:46:1f"`
	IP             string     `json:"ip" doc:"IP address of the DHCP lease." example:"192.168.1.22"`
	Name           string     `json:"name,omitempty" doc:"Name the resolver found for the device. Absent when it has none." example:"iPhone.lan"`
	LeaseExpiresAt time.Time  `json:"leaseExpiresAt" doc:"When the DHCP lease expires, RFC 3339 in UTC." example:"2026-11-01T10:00:00Z"`
	Queries        int64      `json:"queries" doc:"DNS queries over the last 24 hours." example:"1250"`
	Blocked        int64      `json:"blocked" doc:"Queries the resolver blocked over the last 24 hours." example:"87"`
	LastQueryAt    *time.Time `json:"lastQueryAt,omitempty" doc:"Last time the device queried, to 5 minutes. Absent when it did not in the last 24 hours." example:"2026-10-31T18:25:00Z"`
}

func newClientBody(c domainnetwork.Client) ClientBody {
	return ClientBody{
		ID:             c.ID,
		MAC:            c.MAC,
		IP:             c.IP,
		Name:           c.Name,
		LeaseExpiresAt: c.LeaseExpiresAt,
		Queries:        c.Queries,
		Blocked:        c.Blocked,
		LastQueryAt:    c.LastQueryAt,
	}
}
