package supervision

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appsupervision "jarvis/server/internal/app/supervision"
	"jarvis/server/internal/httpapi/problem"
)

type listStoragesInput struct {
	Node string `path:"node" doc:"Proxmox node name." example:"pve"`
}

type listStoragesOutput struct {
	Body struct {
		Items []storageBody `json:"items" nullable:"false" doc:"Storages of the node by name."`
	}
}

func registerListStorages(api huma.API, problems *problem.Mapper, hardware *appsupervision.Hardware) {
	huma.Register(api, huma.Operation{
		OperationID: "listStorages",
		Method:      http.MethodGet,
		Path:        "/nodes/{node}/storages",
		Summary:     "List the storages of a node",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *listStoragesInput) (*listStoragesOutput, error) {
		storages, err := hardware.Storages(ctx, in.Node)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listStoragesOutput{}
		out.Body.Items = make([]storageBody, 0, len(storages))
		for _, s := range storages {
			out.Body.Items = append(out.Body.Items, newStorageBody(s))
		}
		return out, nil
	})
}
