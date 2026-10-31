package network

import (
	"math"

	"jarvis/server/internal/domain/network"
)

// The resolver labels a client by name when it has one, else by IP, so both are tried.
func join(leases []network.Lease, names map[string]string, usage map[string]network.Usage) []network.Client {
	clients := make([]network.Client, 0, len(leases))
	for _, lease := range leases {
		name := names[lease.IP]
		if name == lease.IP {
			name = ""
		}
		client := network.Client{
			ID:             network.ClientID(lease.MAC),
			MAC:            lease.MAC,
			IP:             lease.IP,
			Name:           name,
			LeaseExpiresAt: lease.ExpiresAt,
		}
		var queries, blocked float64
		for _, key := range labelsOf(lease.IP, name) {
			u := usage[key]
			queries += u.Queries
			blocked += u.Blocked
			if u.LastQueryAt != nil && (client.LastQueryAt == nil || u.LastQueryAt.After(*client.LastQueryAt)) {
				client.LastQueryAt = u.LastQueryAt
			}
		}
		client.Queries, client.Blocked = int64(math.Round(queries)), int64(math.Round(blocked))
		clients = append(clients, client)
	}
	return clients
}

func labelsOf(ip, name string) []string {
	if name == "" {
		return []string{ip}
	}
	return []string{ip, name}
}
