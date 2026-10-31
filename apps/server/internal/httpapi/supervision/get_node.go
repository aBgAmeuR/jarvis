package supervision

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appsupervision "jarvis/server/internal/app/supervision"
	"jarvis/server/internal/httpapi/problem"
)

type getNodeInput struct {
	Node string `path:"node" doc:"Proxmox node name." example:"pve"`
}

type getNodeOutput struct {
	Body nodeBody
}

func registerGetNode(api huma.API, problems *problem.Mapper, nodes *appsupervision.Nodes) {
	huma.Register(api, huma.Operation{
		OperationID: "getNode",
		Method:      http.MethodGet,
		Path:        "/nodes/{node}",
		Summary:     "Get a node",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *getNodeInput) (*getNodeOutput, error) {
		node, err := nodes.Get(ctx, in.Node)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getNodeOutput{Body: newNodeBody(node)}, nil
	})
}
