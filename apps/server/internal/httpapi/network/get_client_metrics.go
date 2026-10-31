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

type getClientMetricsInput struct {
	dto.WindowInput
	Client string   `path:"mac" doc:"MAC address in lowercase, with hyphens." example:"bc-07-1d-35-46-1f"`
	Series []string `query:"series" explode:"false" enum:"queriesPerSecond,blockedPerSecond" doc:"Series to read, comma-separated. All when absent." example:"queriesPerSecond"`
}

type getClientMetricsOutput struct {
	Body dto.MetricsBody
}

func registerGetClientMetrics(api huma.API, problems *problem.Mapper, clients *appnetwork.Clients) {
	huma.Register(api, huma.Operation{
		OperationID: "getClientMetrics",
		Method:      http.MethodGet,
		Path:        "/clients/{mac}/metrics",
		Summary:     "Get the DNS activity of a client",
		Tags:        []string{"network"},
	}, func(ctx context.Context, in *getClientMetricsInput) (*getClientMetricsOutput, error) {
		names := make([]domainnetwork.ClientSeries, 0, len(in.Series))
		for _, s := range in.Series {
			names = append(names, domainnetwork.ClientSeries(s))
		}
		ts, err := clients.Metrics(ctx, in.Client, names, in.Start, in.End)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getClientMetricsOutput{Body: dto.NewMetricsBody(ts)}, nil
	})
}
