package supervision

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appsupervision "jarvis/server/internal/app/supervision"
	domainsupervision "jarvis/server/internal/domain/supervision"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type getNodeMetricsInput struct {
	dto.WindowInput
	Node   string   `path:"node" doc:"Proxmox node name." example:"pve"`
	Series []string `query:"series" explode:"false" enum:"cpuUsageRatio,memoryUsedBytes,diskUsedBytes,load1,cpuTemperatureCelsius" doc:"Series to read, comma-separated. All when absent." example:"cpuUsageRatio,load1"`
}

type getNodeMetricsOutput struct {
	Body dto.MetricsBody
}

func registerGetNodeMetrics(api huma.API, problems *problem.Mapper, metrics *appsupervision.Metrics) {
	huma.Register(api, huma.Operation{
		OperationID: "getNodeMetrics",
		Method:      http.MethodGet,
		Path:        "/nodes/{node}/metrics",
		Summary:     "Get the metrics of a node",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *getNodeMetricsInput) (*getNodeMetricsOutput, error) {
		names := make([]domainsupervision.NodeSeries, 0, len(in.Series))
		for _, s := range in.Series {
			names = append(names, domainsupervision.NodeSeries(s))
		}
		ts, err := metrics.Node(ctx, in.Node, names, in.Start, in.End)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getNodeMetricsOutput{Body: dto.NewMetricsBody(ts)}, nil
	})
}
