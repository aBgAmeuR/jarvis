package maintenance

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appmaintenance "jarvis/server/internal/app/maintenance"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type listDriftItemsOutput struct {
	Body struct {
		Items    []driftItemBody    `json:"items" nullable:"false" doc:"Differences between Proxmox and homelab-iac."`
		Degraded []dto.DegradedBody `json:"degraded" nullable:"false" doc:"Sources that failed. The other differences are still listed."`
	}
}

func registerListDriftItems(api huma.API, problems *problem.Mapper, drift *appmaintenance.Drift) {
	huma.Register(api, huma.Operation{
		OperationID: "listDriftItems",
		Method:      http.MethodGet,
		Path:        "/drift-items",
		Summary:     "List the differences between Proxmox and homelab-iac",
		Tags:        []string{"maintenance"},
	}, func(ctx context.Context, _ *struct{}) (*listDriftItemsOutput, error) {
		list, err := drift.List(ctx)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listDriftItemsOutput{}
		out.Body.Items = make([]driftItemBody, 0, len(list.Items))
		for _, i := range list.Items {
			out.Body.Items = append(out.Body.Items, newDriftItemBody(i))
		}
		out.Body.Degraded = dto.NewDegradedBodies(list.Degraded)
		return out, nil
	})
}
