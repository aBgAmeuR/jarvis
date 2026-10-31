package network

import "time"

type Peer struct {
	Name       string
	Interface  string
	AllowedIPs string
	// LastHandshakeAt is nil when the peer never connected.
	LastHandshakeAt  *time.Time
	ReceivedBytes    uint64
	TransmittedBytes uint64
}

type PeerSeries string

const (
	PeerReceivePerSecond  PeerSeries = "receiveBytesPerSecond"
	PeerTransmitPerSecond PeerSeries = "transmitBytesPerSecond"
)

func AllPeerSeries() []PeerSeries {
	return []PeerSeries{PeerReceivePerSecond, PeerTransmitPerSecond}
}
