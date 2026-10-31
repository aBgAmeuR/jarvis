package network

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appnetwork "jarvis/server/internal/app/network"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type ClientDetailBody struct {
	ClientBody
	Degraded []dto.DegradedBody `json:"degraded" nullable:"false" doc:"Sources that failed. The client is still returned."`
}

type getClientInput struct {
	Client string `path:"mac" doc:"MAC address in lowercase, with hyphens." example:"bc-07-1d-35-46-1f"`
}

type getClientOutput struct {
	Body ClientDetailBody
}

func registerGetClient(api huma.API, problems *problem.Mapper, clients *appnetwork.Clients) {
	huma.Register(api, huma.Operation{
		OperationID: "getClient",
		Method:      http.MethodGet,
		Path:        "/clients/{mac}",
		Summary:     "Get a network client",
		Tags:        []string{"network"},
	}, func(ctx context.Context, in *getClientInput) (*getClientOutput, error) {
		client, degraded, err := clients.Get(ctx, in.Client)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getClientOutput{Body: ClientDetailBody{ClientBody: newClientBody(client), Degraded: dto.NewDegradedBodies(degraded)}}, nil
	})
}
