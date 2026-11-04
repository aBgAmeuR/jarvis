package httpapi

import (
	"log/slog"

	apphealth "jarvis/server/internal/app/health"
	appidentity "jarvis/server/internal/app/identity"
	applogs "jarvis/server/internal/app/logs"
	appoverview "jarvis/server/internal/app/overview"
	appsearch "jarvis/server/internal/app/search"
	appservices "jarvis/server/internal/app/services"
	"jarvis/server/internal/domain"
	"jarvis/server/internal/httpapi/availability"
	"jarvis/server/internal/httpapi/maintenance"
	"jarvis/server/internal/httpapi/network"
	"jarvis/server/internal/httpapi/supervision"
)

// The OpenAPI generator leaves the services nil: no route runs.
type Deps struct {
	Logger       *slog.Logger
	Meta         domain.Meta
	Grafana      domain.Grafana
	Identity     *appidentity.Resolver
	Health       *apphealth.Service
	Search       *appsearch.Service
	Overview     *appoverview.Service
	Logs         *applogs.Logs
	Services     *appservices.Services
	Supervision  supervision.Services
	Availability availability.Services
	Network      network.Services
	Maintenance  maintenance.Services
}
