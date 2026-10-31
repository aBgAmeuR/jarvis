package httpapi

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"jarvis/server/internal/domain"
)

type metaBody struct {
	Version   string `json:"version" doc:"Release version." example:"1.2.3"`
	Commit    string `json:"commit" doc:"Git commit the binary was built from." example:"d13fb0c"`
	BuildDate string `json:"buildDate" doc:"Build date, RFC 3339 in UTC." example:"2026-10-31T10:00:00Z"`
}

type metaOutput struct {
	Body metaBody
}

func newMetaOutput(m domain.Meta) *metaOutput {
	return &metaOutput{Body: metaBody{Version: m.Version, Commit: m.Commit, BuildDate: m.BuildDate}}
}

func registerMeta(api huma.API, meta domain.Meta) {
	huma.Register(api, huma.Operation{
		OperationID: "getMeta",
		Method:      http.MethodGet,
		Path:        "/meta",
		Summary:     "Get build information",
		Tags:        []string{"common"},
	}, func(context.Context, *struct{}) (*metaOutput, error) {
		return newMetaOutput(meta), nil
	})
}
