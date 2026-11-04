package logs

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"

	applogs "jarvis/server/internal/app/logs"
	domainlogs "jarvis/server/internal/domain/logs"
	"jarvis/server/internal/httpapi/problem"
)

type volumeLogsInput struct {
	Guest   int       `query:"guest" doc:"Only lines of this guest VMID." example:"116"`
	Service string    `query:"service" doc:"Only lines of this unit or container." example:"blocky"`
	Level   string    `query:"level" enum:"emerg,alert,crit,error,warning,notice,info,debug" doc:"Only lines of this journal priority."`
	Q       string    `query:"q" doc:"Only lines containing this text, case-sensitive." example:"timeout"`
	Start   time.Time `query:"start" doc:"Start of the window, RFC 3339. Defaults to 24 hours before end." example:"2026-10-30T10:00:00Z"`
	End     time.Time `query:"end" doc:"End of the window, RFC 3339. Defaults to now." example:"2026-10-31T10:00:00Z"`
}

type volumeLogsOutput struct {
	Body struct {
		Start       time.Time         `json:"start" doc:"Start of the window actually read, aligned on its step." example:"2026-10-30T10:00:00Z"`
		End         time.Time         `json:"end" doc:"End of the window actually read, aligned on its step." example:"2026-10-31T10:00:00Z"`
		StepSeconds int               `json:"stepSeconds" doc:"Width of a bucket." example:"120"`
		Items       []volumeCountBody `json:"items" nullable:"false" doc:"Lines per bucket and level. Empty buckets are left out."`
	}
}

func registerVolumeLogs(api huma.API, problems *problem.Mapper, logs *applogs.Logs) {
	huma.Register(api, huma.Operation{
		OperationID: "getLogVolume",
		Method:      http.MethodGet,
		Path:        "/logs/volume",
		Summary:     "Count logs over time",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *volumeLogsInput) (*volumeLogsOutput, error) {
		volume, err := logs.Volume(ctx, domainlogs.Query{
			VMID: in.Guest, Service: in.Service, Level: in.Level, Contains: in.Q,
			Start: in.Start, End: in.End,
		})
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &volumeLogsOutput{}
		out.Body.Start = volume.Window.Start
		out.Body.End = volume.Window.End
		out.Body.StepSeconds = int(volume.Window.Step.Seconds())
		out.Body.Items = make([]volumeCountBody, 0, len(volume.Counts))
		for _, c := range volume.Counts {
			out.Body.Items = append(out.Body.Items, newVolumeCountBody(c))
		}
		return out, nil
	})
}
