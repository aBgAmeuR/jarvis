package search

import (
	"github.com/danielgtaylor/huma/v2"

	appsearch "jarvis/server/internal/app/search"
	"jarvis/server/internal/httpapi/problem"
)

func Register(api huma.API, problems *problem.Mapper, search *appsearch.Service) {
	registerSearch(api, problems, search)
}
