package network

import (
	"time"

	domainnetwork "jarvis/server/internal/domain/network"
)

type peerBody struct {
	Name             string     `json:"name" doc:"Peer name." example:"client-1"`
	Interface        string     `json:"interface" doc:"WireGuard interface." example:"wg0"`
	AllowedIPs       string     `json:"allowedIps" doc:"Addresses the peer may use." example:"10.0.0.2/32"`
	LastHandshakeAt  *time.Time `json:"lastHandshakeAt,omitempty" doc:"Last handshake, RFC 3339 in UTC. Absent when the peer never connected." example:"2026-10-31T18:20:00Z"`
	ReceivedBytes    uint64     `json:"receivedBytes" doc:"Bytes received from the peer since the interface came up." example:"1048576"`
	TransmittedBytes uint64     `json:"transmittedBytes" doc:"Bytes sent to the peer since the interface came up." example:"2097152"`
}

func newPeerBody(p domainnetwork.Peer) peerBody {
	return peerBody{
		Name:             p.Name,
		Interface:        p.Interface,
		AllowedIPs:       p.AllowedIPs,
		LastHandshakeAt:  p.LastHandshakeAt,
		ReceivedBytes:    p.ReceivedBytes,
		TransmittedBytes: p.TransmittedBytes,
	}
}
