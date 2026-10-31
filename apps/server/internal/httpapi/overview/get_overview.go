package overview

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appoverview "jarvis/server/internal/app/overview"
	"jarvis/server/internal/httpapi/dto"
)

type getOverviewOutput struct {
	Body overviewBody
}

func registerGetOverview(api huma.API, overview *appoverview.Service) {
	huma.Register(api, huma.Operation{
		OperationID: "getOverview",
		Method:      http.MethodGet,
		Path:        "/overview",
		Summary:     "Get the overview",
		Description: "Global status, node resources, key counters and the alerts to handle, in one call.",
		Tags:        []string{"overview"},
	}, func(ctx context.Context, _ *struct{}) (*getOverviewOutput, error) {
		result := overview.Get(ctx)
		return &getOverviewOutput{Body: newOverviewBody(result.Overview, dto.NewDegradedBodies(result.Degraded))}, nil
	})
}
