package availability

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	domainavailability "jarvis/server/internal/domain/availability"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type getProbeMetricsInput struct {
	dto.WindowInput
	Probe  string   `path:"probe" doc:"Probe id." example:"icmp-116"`
	Series []string `query:"series" explode:"false" enum:"durationSeconds,success" doc:"Series to read, comma-separated. All when absent." example:"durationSeconds"`
}

type getProbeMetricsOutput struct {
	Body dto.MetricsBody
}

func registerGetProbeMetrics(api huma.API, problems *problem.Mapper, probes *appavailability.Probes) {
	huma.Register(api, huma.Operation{
		OperationID: "getProbeMetrics",
		Method:      http.MethodGet,
		Path:        "/probes/{probe}/metrics",
		Summary:     "Get the metrics of a probe",
		Tags:        []string{"availability"},
	}, func(ctx context.Context, in *getProbeMetricsInput) (*getProbeMetricsOutput, error) {
		names := make([]domainavailability.ProbeSeries, 0, len(in.Series))
		for _, s := range in.Series {
			names = append(names, domainavailability.ProbeSeries(s))
		}
		ts, err := probes.Metrics(ctx, in.Probe, names, in.Start, in.End)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getProbeMetricsOutput{Body: dto.NewMetricsBody(ts)}, nil
	})
}
