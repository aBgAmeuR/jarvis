package network

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appnetwork "jarvis/server/internal/app/network"
	domainnetwork "jarvis/server/internal/domain/network"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type getPeerMetricsInput struct {
	dto.WindowInput
	Peer   string   `path:"peer" doc:"Peer name." example:"client-1"`
	Series []string `query:"series" explode:"false" enum:"receiveBytesPerSecond,transmitBytesPerSecond" doc:"Series to read, comma-separated. All when absent." example:"receiveBytesPerSecond"`
}

type getPeerMetricsOutput struct {
	Body dto.MetricsBody
}

func registerGetPeerMetrics(api huma.API, problems *problem.Mapper, peers *appnetwork.Peers) {
	huma.Register(api, huma.Operation{
		OperationID: "getVpnPeerMetrics",
		Method:      http.MethodGet,
		Path:        "/vpn-peers/{peer}/metrics",
		Summary:     "Get the traffic of a VPN peer",
		Tags:        []string{"network"},
	}, func(ctx context.Context, in *getPeerMetricsInput) (*getPeerMetricsOutput, error) {
		names := make([]domainnetwork.PeerSeries, 0, len(in.Series))
		for _, s := range in.Series {
			names = append(names, domainnetwork.PeerSeries(s))
		}
		ts, err := peers.Metrics(ctx, in.Peer, names, in.Start, in.End)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getPeerMetricsOutput{Body: dto.NewMetricsBody(ts)}, nil
	})
}
