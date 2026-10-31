package network

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appnetwork "jarvis/server/internal/app/network"
	"jarvis/server/internal/httpapi/problem"
)

type getPeerInput struct {
	Peer string `path:"peer" doc:"Peer name." example:"client-1"`
}

type getPeerOutput struct {
	Body peerBody
}

func registerGetPeer(api huma.API, problems *problem.Mapper, peers *appnetwork.Peers) {
	huma.Register(api, huma.Operation{
		OperationID: "getVpnPeer",
		Method:      http.MethodGet,
		Path:        "/vpn-peers/{peer}",
		Summary:     "Get a VPN peer",
		Tags:        []string{"network"},
	}, func(ctx context.Context, in *getPeerInput) (*getPeerOutput, error) {
		peer, err := peers.Get(ctx, in.Peer)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getPeerOutput{Body: newPeerBody(peer)}, nil
	})
}
