package services

import (
	"github.com/danielgtaylor/huma/v2"

	appservices "jarvis/server/internal/app/services"
	"jarvis/server/internal/httpapi/problem"
)

func Register(api huma.API, problems *problem.Mapper, services *appservices.Services) {
	registerListServices(api, problems, services)
	registerGetService(api, problems, services)
	registerGetServiceMetrics(api, problems, services)
}
