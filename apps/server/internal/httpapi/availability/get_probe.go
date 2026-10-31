package availability

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	"jarvis/server/internal/httpapi/problem"
)

type getProbeInput struct {
	Probe string `path:"probe" doc:"Probe id." example:"icmp-116"`
}

type getProbeOutput struct {
	Body probeBody
}

func registerGetProbe(api huma.API, problems *problem.Mapper, probes *appavailability.Probes) {
	huma.Register(api, huma.Operation{
		OperationID: "getProbe",
		Method:      http.MethodGet,
		Path:        "/probes/{probe}",
		Summary:     "Get a probe",
		Tags:        []string{"availability"},
	}, func(ctx context.Context, in *getProbeInput) (*getProbeOutput, error) {
		probe, err := probes.Get(ctx, in.Probe)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getProbeOutput{Body: newProbeBody(probe)}, nil
	})
}
