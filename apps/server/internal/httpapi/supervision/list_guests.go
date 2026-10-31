package supervision

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appsupervision "jarvis/server/internal/app/supervision"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type listGuestsOutput struct {
	Body struct {
		Items    []GuestBody        `json:"items" nullable:"false" doc:"Guests by VMID."`
		Degraded []dto.DegradedBody `json:"degraded" nullable:"false" doc:"Sources that failed. The guests are still listed."`
	}
}

func registerListGuests(api huma.API, problems *problem.Mapper, guests *appsupervision.Guests) {
	huma.Register(api, huma.Operation{
		OperationID: "listGuests",
		Method:      http.MethodGet,
		Path:        "/guests",
		Summary:     "List guests",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, _ *struct{}) (*listGuestsOutput, error) {
		list, err := guests.List(ctx)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listGuestsOutput{}
		out.Body.Items = make([]GuestBody, 0, len(list.Items))
		for _, g := range list.Items {
			out.Body.Items = append(out.Body.Items, newGuestBody(g))
		}
		out.Body.Degraded = dto.NewDegradedBodies(list.Degraded)
		return out, nil
	})
}
