package network

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appnetwork "jarvis/server/internal/app/network"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type listClientsInput struct {
	Sort  string `query:"sort" enum:"name,-name,queries,-queries,blocked,-blocked,lastQueryAt,-lastQueryAt" default:"name" doc:"Sort field, prefixed with - for descending order."`
	Limit int    `query:"limit" minimum:"1" maximum:"1000" doc:"Maximum number of clients. All when absent." example:"10"`
}

type listClientsOutput struct {
	Body struct {
		Items    []ClientBody       `json:"items" nullable:"false" doc:"Devices that hold a DHCP lease."`
		Degraded []dto.DegradedBody `json:"degraded" nullable:"false" doc:"Sources that failed. The clients are still listed, without their counts or names."`
	}
}

func registerListClients(api huma.API, problems *problem.Mapper, clients *appnetwork.Clients) {
	huma.Register(api, huma.Operation{
		OperationID: "listClients",
		Method:      http.MethodGet,
		Path:        "/clients",
		Summary:     "List network clients",
		Tags:        []string{"network"},
	}, func(ctx context.Context, in *listClientsInput) (*listClientsOutput, error) {
		list, err := clients.List(ctx, appnetwork.ClientSort(in.Sort), in.Limit)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listClientsOutput{}
		out.Body.Items = make([]ClientBody, 0, len(list.Items))
		for _, c := range list.Items {
			out.Body.Items = append(out.Body.Items, newClientBody(c))
		}
		out.Body.Degraded = dto.NewDegradedBodies(list.Degraded)
		return out, nil
	})
}
