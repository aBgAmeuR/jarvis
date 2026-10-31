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

type getGuestMetricsInput struct {
	dto.WindowInput
	VMID   int      `path:"vmid" doc:"Proxmox guest id." example:"116"`
	Series []string `query:"series" explode:"false" enum:"cpuUsageRatio,memoryUsedBytes,diskUsedBytes,diskReadBytesPerSecond,diskWriteBytesPerSecond,networkReceiveBytesPerSecond,networkTransmitBytesPerSecond" doc:"Series to read, comma-separated. All when absent." example:"cpuUsageRatio,memoryUsedBytes"`
}

type getGuestMetricsOutput struct {
	Body dto.MetricsBody
}

func registerGetGuestMetrics(api huma.API, problems *problem.Mapper, metrics *appsupervision.Metrics) {
	huma.Register(api, huma.Operation{
		OperationID: "getGuestMetrics",
		Method:      http.MethodGet,
		Path:        "/guests/{vmid}/metrics",
		Summary:     "Get the metrics of a guest",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *getGuestMetricsInput) (*getGuestMetricsOutput, error) {
		names := make([]domainsupervision.GuestSeries, 0, len(in.Series))
		for _, s := range in.Series {
			names = append(names, domainsupervision.GuestSeries(s))
		}
		ts, err := metrics.Guest(ctx, in.VMID, names, in.Start, in.End)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getGuestMetricsOutput{Body: dto.NewMetricsBody(ts)}, nil
	})
}
