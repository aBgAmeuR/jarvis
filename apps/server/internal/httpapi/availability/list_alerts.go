package availability

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	"jarvis/server/internal/domain"
	domainavailability "jarvis/server/internal/domain/availability"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type listAlertsInput struct {
	dto.WindowInput
	State    string `query:"state" enum:"firing,suppressed,resolved" default:"firing" doc:"Alert state. Resolved alerts are rebuilt from Prometheus over the window, which the other states ignore."`
	Severity string `query:"severity" enum:"critical,warning,info" doc:"Only alerts of this severity."`
	Guest    string `query:"guest" doc:"Only alerts of this guest VMID." example:"116"`
}

type listAlertsOutput struct {
	Body struct {
		Items []alertBody `json:"items" nullable:"false" doc:"Alerts, most severe first, then most recent first. Resolved alerts come most recently ended first."`
	}
}

func registerListAlerts(api huma.API, problems *problem.Mapper, alerts *appavailability.Alerts) {
	huma.Register(api, huma.Operation{
		OperationID: "listAlerts",
		Method:      http.MethodGet,
		Path:        "/alerts",
		Summary:     "List alerts",
		Tags:        []string{"availability"},
	}, func(ctx context.Context, in *listAlertsInput) (*listAlertsOutput, error) {
		filter := appavailability.AlertFilter{
			State:    domainavailability.AlertState(in.State),
			Severity: domain.Severity(in.Severity),
			Guest:    in.Guest,
		}
		var (
			list []domainavailability.Alert
			err  error
		)
		if filter.State == domainavailability.AlertResolved {
			list, err = alerts.Resolved(ctx, filter, in.Start, in.End)
		} else {
			list, err = alerts.List(ctx, filter)
		}
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listAlertsOutput{}
		out.Body.Items = make([]alertBody, 0, len(list))
		for _, a := range list {
			out.Body.Items = append(out.Body.Items, newAlertBody(a))
		}
		return out, nil
	})
}
