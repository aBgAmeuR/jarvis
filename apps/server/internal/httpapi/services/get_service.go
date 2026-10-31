package services

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appservices "jarvis/server/internal/app/services"
	"jarvis/server/internal/httpapi/problem"
)

type getServiceInput struct {
	Service string `path:"service" doc:"Service slug." example:"grafana"`
}

type getServiceOutput struct {
	Body ServiceBody
}

func registerGetService(api huma.API, problems *problem.Mapper, services *appservices.Services) {
	huma.Register(api, huma.Operation{
		OperationID: "getService",
		Method:      http.MethodGet,
		Path:        "/services/{service}",
		Summary:     "Get a service",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *getServiceInput) (*getServiceOutput, error) {
		service, err := services.Get(ctx, in.Service)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getServiceOutput{Body: newServiceBody(service)}, nil
	})
}
