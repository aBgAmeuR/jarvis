package supervision

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appsupervision "jarvis/server/internal/app/supervision"
	"jarvis/server/internal/httpapi/problem"
)

type listZPoolsInput struct {
	Node string `path:"node" doc:"Proxmox node name." example:"pve"`
}

type listZPoolsOutput struct {
	Body struct {
		Items []zpoolBody `json:"items" nullable:"false" doc:"ZFS pools of the node by name."`
	}
}

func registerListZPools(api huma.API, problems *problem.Mapper, hardware *appsupervision.Hardware) {
	huma.Register(api, huma.Operation{
		OperationID: "listZpools",
		Method:      http.MethodGet,
		Path:        "/nodes/{node}/zpools",
		Summary:     "List the ZFS pools of a node",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *listZPoolsInput) (*listZPoolsOutput, error) {
		pools, err := hardware.ZPools(ctx, in.Node)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listZPoolsOutput{}
		out.Body.Items = make([]zpoolBody, 0, len(pools))
		for _, p := range pools {
			out.Body.Items = append(out.Body.Items, newZPoolBody(p))
		}
		return out, nil
	})
}
