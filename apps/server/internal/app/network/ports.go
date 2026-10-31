package network

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

type ClientSource interface {
	Leases(ctx context.Context) ([]network.Lease, error)
	// ClientUsage is keyed by the client label of the resolver: a name when it knows one, else an IP.
	ClientUsage(ctx context.Context) (map[string]network.Usage, error)
	ClientSeries(ctx context.Context, keys []string, names []network.ClientSeries, w domain.Window) ([]domain.Series, error)
}

type NameSource interface {
	ClientNames(ctx context.Context, w domain.Window) (map[string]string, error)
}

type DomainSource interface {
	ClientDomains(ctx context.Context, ip string, status network.DomainStatus, w domain.Window, limit int) ([]network.DomainCount, error)
	DNSDomains(ctx context.Context, status network.DomainStatus, w domain.Window, limit int) ([]network.DomainCount, error)
}

type DNSSource interface {
	DNS(ctx context.Context) (network.DNS, error)
	DNSSeries(ctx context.Context, names []network.DNSSeries, w domain.Window) ([]domain.Series, error)
}

type PeerSource interface {
	Peers(ctx context.Context) ([]network.Peer, error)
	PeerSeries(ctx context.Context, peer string, names []network.PeerSeries, w domain.Window) ([]domain.Series, error)
}
