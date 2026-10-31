package network

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appnetwork "jarvis/server/internal/app/network"
	"jarvis/server/internal/httpapi/problem"
)

type listPeersOutput struct {
	Body struct {
		Items []peerBody `json:"items" nullable:"false" doc:"VPN peers by name."`
	}
}

func registerListPeers(api huma.API, problems *problem.Mapper, peers *appnetwork.Peers) {
	huma.Register(api, huma.Operation{
		OperationID: "listVpnPeers",
		Method:      http.MethodGet,
		Path:        "/vpn-peers",
		Summary:     "List VPN peers",
		Tags:        []string{"network"},
	}, func(ctx context.Context, _ *struct{}) (*listPeersOutput, error) {
		list, err := peers.List(ctx)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listPeersOutput{}
		out.Body.Items = make([]peerBody, 0, len(list))
		for _, p := range list {
			out.Body.Items = append(out.Body.Items, newPeerBody(p))
		}
		return out, nil
	})
}
