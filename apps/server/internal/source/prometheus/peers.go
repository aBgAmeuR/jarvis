package prometheus

import (
	"context"
	"slices"
	"strings"
	"sync"
	"time"

	"jarvis/server/internal/domain/network"
	"jarvis/server/internal/source/failure"
)

func (c *Client) Peers(ctx context.Context) ([]network.Peer, error) {
	exprs := [...]string{
		"wireguard_peer_last_handshake_seconds",
		"wireguard_peer_receive_bytes_total",
		"wireguard_peer_transmit_bytes_total",
	}
	var results [len(exprs)][]entry
	var errs [len(exprs)]error
	var wg sync.WaitGroup
	for i, expr := range exprs {
		wg.Go(func() { results[i], errs[i] = c.queryVector(ctx, expr) })
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return nil, failure.Wrap(name, err)
		}
	}

	received, transmitted := map[string]float64{}, map[string]float64{}
	for _, e := range results[1] {
		received[e.labels["peer"]] = e.value
	}
	for _, e := range results[2] {
		transmitted[e.labels["peer"]] = e.value
	}
	peers := make([]network.Peer, 0, len(results[0]))
	for _, e := range results[0] {
		peer := network.Peer{
			Name:             e.labels["peer"],
			Interface:        e.labels["interface"],
			AllowedIPs:       e.labels["allowed_ips"],
			ReceivedBytes:    uint64(received[e.labels["peer"]]),
			TransmittedBytes: uint64(transmitted[e.labels["peer"]]),
		}
		if e.value > 0 {
			at := time.Unix(int64(e.value), 0).UTC()
			peer.LastHandshakeAt = &at
		}
		peers = append(peers, peer)
	}
	slices.SortFunc(peers, func(a, b network.Peer) int { return strings.Compare(a.Name, b.Name) })
	return peers, nil
}
