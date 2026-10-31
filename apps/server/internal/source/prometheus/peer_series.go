package prometheus

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

func (c *Client) PeerSeries(ctx context.Context, peer string, names []network.PeerSeries, w domain.Window) ([]domain.Series, error) {
	selector := "{peer=" + quote(peer) + "}"
	rate := rateRange(w)
	queries := make([]query, 0, len(names))
	for _, n := range names {
		q := query{name: string(n), unit: domain.UnitBytesPerSecond}
		switch n {
		case network.PeerReceivePerSecond:
			q.expr = "max(rate(wireguard_peer_receive_bytes_total" + selector + rate + "))"
		case network.PeerTransmitPerSecond:
			q.expr = "max(rate(wireguard_peer_transmit_bytes_total" + selector + rate + "))"
		default:
			continue
		}
		queries = append(queries, q)
	}
	return c.runQueries(ctx, queries, w)
}
