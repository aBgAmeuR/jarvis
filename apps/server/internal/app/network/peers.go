package network

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

type Peers struct {
	source PeerSource
	now    func() time.Time
}

func NewPeers(source PeerSource) *Peers {
	return &Peers{source: source, now: time.Now}
}

func (p *Peers) List(ctx context.Context) ([]network.Peer, error) {
	all, err := p.source.Peers(ctx)
	if err != nil {
		return nil, fmt.Errorf("list vpn peers: %w", err)
	}
	peers := slices.Clone(all)
	slices.SortFunc(peers, func(a, b network.Peer) int { return cmp.Compare(a.Name, b.Name) })
	return peers, nil
}

func (p *Peers) Get(ctx context.Context, name string) (network.Peer, error) {
	peers, err := p.List(ctx)
	if err != nil {
		return network.Peer{}, err
	}
	for _, peer := range peers {
		if peer.Name == name {
			return peer, nil
		}
	}
	return network.Peer{}, domain.NewNotFound("VPN peer", name)
}

func (p *Peers) Metrics(ctx context.Context, name string, names []network.PeerSeries, start, end time.Time) (domain.TimeSeries, error) {
	w, err := domain.NewWindow(start, end, p.now())
	if err != nil {
		return domain.TimeSeries{}, err
	}
	if _, err := p.Get(ctx, name); err != nil {
		return domain.TimeSeries{}, err
	}
	if len(names) == 0 {
		names = network.AllPeerSeries()
	}
	series, err := p.source.PeerSeries(ctx, name, slices.Compact(slices.Sorted(slices.Values(names))), w)
	if err != nil {
		return domain.TimeSeries{}, fmt.Errorf("vpn peer %s metrics: %w", name, err)
	}
	return domain.TimeSeries{Window: w, Series: series}, nil
}
