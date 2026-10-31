package services

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appservices "jarvis/server/internal/app/services"
	domainservices "jarvis/server/internal/domain/services"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type getServiceMetricsInput struct {
	dto.WindowInput
	Service string   `path:"service" doc:"Service slug." example:"grafana"`
	Series  []string `query:"series" explode:"false" enum:"requestsPerSecond,errorRatio,latencyP95Seconds" doc:"Series to read, comma-separated. All when absent." example:"requestsPerSecond"`
}

type getServiceMetricsOutput struct {
	Body dto.MetricsBody
}

func registerGetServiceMetrics(api huma.API, problems *problem.Mapper, services *appservices.Services) {
	huma.Register(api, huma.Operation{
		OperationID: "getServiceMetrics",
		Method:      http.MethodGet,
		Path:        "/services/{service}/metrics",
		Summary:     "Get the request metrics of a service",
		Description: "The series are empty when the service is not instrumented.",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *getServiceMetricsInput) (*getServiceMetricsOutput, error) {
		names := make([]domainservices.Series, 0, len(in.Series))
		for _, s := range in.Series {
			names = append(names, domainservices.Series(s))
		}
		ts, err := services.Metrics(ctx, in.Service, names, in.Start, in.End)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getServiceMetricsOutput{Body: dto.NewMetricsBody(ts)}, nil
	})
}
