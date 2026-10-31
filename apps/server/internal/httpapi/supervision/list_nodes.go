package supervision

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appsupervision "jarvis/server/internal/app/supervision"
	"jarvis/server/internal/httpapi/problem"
)

type listNodesOutput struct {
	Body struct {
		Items []nodeBody `json:"items" nullable:"false" doc:"Proxmox nodes."`
	}
}

func registerListNodes(api huma.API, problems *problem.Mapper, nodes *appsupervision.Nodes) {
	huma.Register(api, huma.Operation{
		OperationID: "listNodes",
		Method:      http.MethodGet,
		Path:        "/nodes",
		Summary:     "List nodes",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, _ *struct{}) (*listNodesOutput, error) {
		list, err := nodes.List(ctx)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listNodesOutput{}
		out.Body.Items = make([]nodeBody, 0, len(list))
		for _, n := range list {
			out.Body.Items = append(out.Body.Items, newNodeBody(n))
		}
		return out, nil
	})
}
