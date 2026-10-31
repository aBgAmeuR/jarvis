package supervision

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appsupervision "jarvis/server/internal/app/supervision"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type GuestDetailBody struct {
	GuestBody
	Degraded []dto.DegradedBody `json:"degraded" nullable:"false" doc:"Sources that failed. The guest is still returned."`
}

type getGuestInput struct {
	VMID int `path:"vmid" doc:"Proxmox guest id." example:"116"`
}

type getGuestOutput struct {
	Body GuestDetailBody
}

func registerGetGuest(api huma.API, problems *problem.Mapper, guests *appsupervision.Guests) {
	huma.Register(api, huma.Operation{
		OperationID: "getGuest",
		Method:      http.MethodGet,
		Path:        "/guests/{vmid}",
		Summary:     "Get a guest",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *getGuestInput) (*getGuestOutput, error) {
		guest, degraded, err := guests.Get(ctx, in.VMID)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getGuestOutput{Body: GuestDetailBody{GuestBody: newGuestBody(guest), Degraded: dto.NewDegradedBodies(degraded)}}, nil
	})
}
