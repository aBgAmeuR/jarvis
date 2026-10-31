package availability

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	"jarvis/server/internal/httpapi/problem"
)

type getSilenceInput struct {
	ID string `path:"id" doc:"Alertmanager silence id." example:"f3b1c0aa-3f3d-4a0e-9a9e-7c9d5d3a1b11"`
}

type getSilenceOutput struct {
	Body silenceBody
}

func registerGetSilence(api huma.API, problems *problem.Mapper, silences *appavailability.Silences) {
	huma.Register(api, huma.Operation{
		OperationID: "getSilence",
		Method:      http.MethodGet,
		Path:        "/silences/{id}",
		Summary:     "Get a silence",
		Tags:        []string{"availability"},
	}, func(ctx context.Context, in *getSilenceInput) (*getSilenceOutput, error) {
		silence, err := silences.Get(ctx, in.ID)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getSilenceOutput{Body: newSilenceBody(silence)}, nil
	})
}
