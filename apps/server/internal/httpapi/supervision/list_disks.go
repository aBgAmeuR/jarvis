package supervision

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appsupervision "jarvis/server/internal/app/supervision"
	"jarvis/server/internal/httpapi/problem"
)

type listDisksInput struct {
	Node string `path:"node" doc:"Proxmox node name." example:"pve"`
}

type listDisksOutput struct {
	Body struct {
		Items []diskBody `json:"items" nullable:"false" doc:"Physical disks of the node by device."`
	}
}

func registerListDisks(api huma.API, problems *problem.Mapper, hardware *appsupervision.Hardware) {
	huma.Register(api, huma.Operation{
		OperationID: "listDisks",
		Method:      http.MethodGet,
		Path:        "/nodes/{node}/disks",
		Summary:     "List the disks of a node with their SMART health",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *listDisksInput) (*listDisksOutput, error) {
		disks, err := hardware.Disks(ctx, in.Node)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listDisksOutput{}
		out.Body.Items = make([]diskBody, 0, len(disks))
		for _, d := range disks {
			out.Body.Items = append(out.Body.Items, newDiskBody(d))
		}
		return out, nil
	})
}
