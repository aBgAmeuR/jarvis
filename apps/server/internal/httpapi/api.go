package httpapi

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"

	"jarvis/server/internal/httpapi/availability"
	"jarvis/server/internal/httpapi/health"
	"jarvis/server/internal/httpapi/identity"
	"jarvis/server/internal/httpapi/logs"
	"jarvis/server/internal/httpapi/maintenance"
	"jarvis/server/internal/httpapi/meta"
	"jarvis/server/internal/httpapi/middleware"
	"jarvis/server/internal/httpapi/network"
	"jarvis/server/internal/httpapi/overview"
	"jarvis/server/internal/httpapi/problem"
	"jarvis/server/internal/httpapi/search"
	"jarvis/server/internal/httpapi/services"
	"jarvis/server/internal/httpapi/supervision"
)

func NewAPI(mux *http.ServeMux, deps Deps) huma.API {
	api := humago.NewWithPrefix(mux, basePath, newConfig(deps.Meta.Version))
	api.UseMiddleware(middleware.Identity(api, deps.Logger, deps.Identity))
	problems := problem.NewMapper(deps.Logger)

	meta.Register(api, deps.Meta, deps.Grafana)
	identity.Register(api)
	health.Register(api, deps.Health)
	search.Register(api, problems, deps.Search)
	overview.Register(api, deps.Overview)
	logs.Register(api, problems, deps.Logs)
	services.Register(api, problems, deps.Services)
	supervision.Register(api, problems, deps.Supervision)
	availability.Register(api, problems, deps.Availability)
	network.Register(api, problems, deps.Network)
	maintenance.Register(api, problems, deps.Maintenance)
	return api
}
