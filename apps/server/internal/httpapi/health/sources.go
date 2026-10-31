package health

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"

	apphealth "jarvis/server/internal/app/health"
	"jarvis/server/internal/httpapi/dto"
)

type SourceBody struct {
	Name      string    `json:"name" doc:"Data source name." example:"prometheus"`
	Up        bool      `json:"up" doc:"Whether the source answered its readiness check." example:"true"`
	Reason    string    `json:"reason,omitempty" doc:"How the source failed. Absent when it is up." example:"Did not answer in time"`
	CheckedAt time.Time `json:"checkedAt" doc:"When the check ran, RFC 3339 in UTC." example:"2026-10-31T10:00:00Z"`
}

type listSourcesOutput struct {
	Body struct {
		Items []SourceBody `json:"items" nullable:"false" doc:"Status of each data source."`
	}
}

func Register(api huma.API, health *apphealth.Service) {
	huma.Register(api, huma.Operation{
		OperationID: "listSources",
		Method:      http.MethodGet,
		Path:        "/sources",
		Summary:     "List data source statuses",
		Tags:        []string{"common"},
	}, func(ctx context.Context, _ *struct{}) (*listSourcesOutput, error) {
		out := &listSourcesOutput{}
		out.Body.Items = []SourceBody{}
		for _, s := range health.Status(ctx) {
			item := SourceBody{Name: s.Name, Up: s.Up, CheckedAt: s.CheckedAt}
			if !s.Up {
				item.Reason = dto.Reason(s.Kind)
			}
			out.Body.Items = append(out.Body.Items, item)
		}
		return out, nil
	})
}
