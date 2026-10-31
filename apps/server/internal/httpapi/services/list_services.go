package services

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appservices "jarvis/server/internal/app/services"
	"jarvis/server/internal/httpapi/problem"
)

type listServicesInput struct {
	Guest int `query:"guest" doc:"Only services of this guest VMID." example:"111"`
}

type listServicesOutput struct {
	Body struct {
		Items []ServiceBody `json:"items" nullable:"false" doc:"Services by slug."`
	}
}

func registerListServices(api huma.API, problems *problem.Mapper, services *appservices.Services) {
	huma.Register(api, huma.Operation{
		OperationID: "listServices",
		Method:      http.MethodGet,
		Path:        "/services",
		Summary:     "List services",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *listServicesInput) (*listServicesOutput, error) {
		list, err := services.List(ctx)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listServicesOutput{}
		out.Body.Items = make([]ServiceBody, 0, len(list))
		for _, s := range list {
			if in.Guest == 0 || s.VMID == in.Guest {
				out.Body.Items = append(out.Body.Items, newServiceBody(s))
			}
		}
		return out, nil
	})
}
