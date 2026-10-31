package availability

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	"jarvis/server/internal/httpapi/problem"
)

type getAlertInput struct {
	Fingerprint string `path:"fingerprint" doc:"Alertmanager fingerprint." example:"3a1f9c0d5b7e2a44"`
}

type getAlertOutput struct {
	Body alertBody
}

func registerGetAlert(api huma.API, problems *problem.Mapper, alerts *appavailability.Alerts) {
	huma.Register(api, huma.Operation{
		OperationID: "getAlert",
		Method:      http.MethodGet,
		Path:        "/alerts/{fingerprint}",
		Summary:     "Get an alert",
		Tags:        []string{"availability"},
	}, func(ctx context.Context, in *getAlertInput) (*getAlertOutput, error) {
		alert, err := alerts.Get(ctx, in.Fingerprint)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getAlertOutput{Body: newAlertBody(alert)}, nil
	})
}
