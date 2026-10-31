package httpapi

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"

	"jarvis/server/internal/domain"
)

func NewAPI(mux *http.ServeMux, meta domain.Meta) huma.API {
	api := humago.NewWithPrefix(mux, basePath, newConfig(meta.Version))
	registerMeta(api, meta)
	return api
}
