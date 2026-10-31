package availability

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	"jarvis/server/internal/httpapi/problem"
)

type listProbesOutput struct {
	Body struct {
		Items []probeBody `json:"items" nullable:"false" doc:"Probes by id."`
	}
}

func registerListProbes(api huma.API, problems *problem.Mapper, probes *appavailability.Probes) {
	huma.Register(api, huma.Operation{
		OperationID: "listProbes",
		Method:      http.MethodGet,
		Path:        "/probes",
		Summary:     "List probes",
		Tags:        []string{"availability"},
	}, func(ctx context.Context, _ *struct{}) (*listProbesOutput, error) {
		list, err := probes.List(ctx)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listProbesOutput{}
		out.Body.Items = make([]probeBody, 0, len(list))
		for _, p := range list {
			out.Body.Items = append(out.Body.Items, newProbeBody(p))
		}
		return out, nil
	})
}
