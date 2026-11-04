package search

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appsearch "jarvis/server/internal/app/search"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type searchInput struct {
	Q     string `query:"q" required:"true" minLength:"1" maxLength:"100" doc:"Text to find in names, ids and details, case-insensitive." example:"graf"`
	Limit int    `query:"limit" minimum:"1" maximum:"50" default:"20" doc:"Maximum number of hits."`
}

type searchOutput struct {
	Body struct {
		Items    []SearchHitBody    `json:"items" nullable:"false" doc:"Hits, best match first."`
		Degraded []dto.DegradedBody `json:"degraded" nullable:"false" doc:"Sources that failed. Their objects are left out."`
	}
}

func registerSearch(api huma.API, problems *problem.Mapper, search *appsearch.Service) {
	huma.Register(api, huma.Operation{
		OperationID: "search",
		Method:      http.MethodGet,
		Path:        "/search",
		Summary:     "Search every object",
		Description: "Guests, services, probes, network clients and VPN peers, for ⌘K.",
		Tags:        []string{"common"},
	}, func(ctx context.Context, in *searchInput) (*searchOutput, error) {
		result, err := search.Search(ctx, in.Q, in.Limit)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &searchOutput{}
		out.Body.Items = make([]SearchHitBody, 0, len(result.Hits))
		for _, h := range result.Hits {
			out.Body.Items = append(out.Body.Items, newSearchHitBody(h))
		}
		out.Body.Degraded = dto.NewDegradedBodies(result.Degraded)
		return out, nil
	})
}
