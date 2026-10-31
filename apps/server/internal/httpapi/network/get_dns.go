package network

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appnetwork "jarvis/server/internal/app/network"
	"jarvis/server/internal/httpapi/problem"
)

type DNSBody struct {
	Queries                int64   `json:"queries" doc:"DNS queries over the last 24 hours." example:"52000"`
	Blocked                int64   `json:"blocked" doc:"Queries blocked over the last 24 hours." example:"2200"`
	Cached                 int64   `json:"cached" doc:"Queries answered from the cache over the last 24 hours." example:"35000"`
	UpstreamLatencySeconds float64 `json:"upstreamLatencySeconds" doc:"Average time the upstream resolver took over the last 24 hours." example:"0.01"`
	QueriesPerSecond       float64 `json:"queriesPerSecond" doc:"Queries per second over the last 5 minutes." example:"0.6"`
}

type getDNSOutput struct {
	Body DNSBody
}

func registerGetDNS(api huma.API, problems *problem.Mapper, dns *appnetwork.DNS) {
	huma.Register(api, huma.Operation{
		OperationID: "getDns",
		Method:      http.MethodGet,
		Path:        "/dns",
		Summary:     "Get the DNS resolver summary",
		Tags:        []string{"network"},
	}, func(ctx context.Context, _ *struct{}) (*getDNSOutput, error) {
		summary, err := dns.Summary(ctx)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getDNSOutput{Body: DNSBody{
			Queries:                summary.Queries,
			Blocked:                summary.Blocked,
			Cached:                 summary.Cached,
			UpstreamLatencySeconds: summary.UpstreamLatencySeconds,
			QueriesPerSecond:       summary.QueriesPerSecond,
		}}, nil
	})
}
