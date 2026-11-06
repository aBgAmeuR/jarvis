package network

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appnetwork "jarvis/server/internal/app/network"
	domainnetwork "jarvis/server/internal/domain/network"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type getDNSMetricsInput struct {
	dto.WindowInput
	Series []string `query:"series" explode:"false" enum:"queriesPerSecond,blockedPerSecond,upstreamLatencySeconds,upstreamLatencyP99Seconds" doc:"Series to read, comma-separated. All when absent." example:"queriesPerSecond"`
}

type getDNSMetricsOutput struct {
	Body dto.MetricsBody
}

func registerGetDNSMetrics(api huma.API, problems *problem.Mapper, dns *appnetwork.DNS) {
	huma.Register(api, huma.Operation{
		OperationID: "getDnsMetrics",
		Method:      http.MethodGet,
		Path:        "/dns/metrics",
		Summary:     "Get the DNS resolver metrics",
		Tags:        []string{"network"},
	}, func(ctx context.Context, in *getDNSMetricsInput) (*getDNSMetricsOutput, error) {
		names := make([]domainnetwork.DNSSeries, 0, len(in.Series))
		for _, s := range in.Series {
			names = append(names, domainnetwork.DNSSeries(s))
		}
		ts, err := dns.Metrics(ctx, names, in.Start, in.End)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getDNSMetricsOutput{Body: dto.NewMetricsBody(ts)}, nil
	})
}
