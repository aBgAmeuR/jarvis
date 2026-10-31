package network

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appnetwork "jarvis/server/internal/app/network"
	domainnetwork "jarvis/server/internal/domain/network"
	"jarvis/server/internal/httpapi/dto"
	"jarvis/server/internal/httpapi/problem"
)

type domainBody struct {
	Domain string `json:"domain" doc:"Domain name, with its trailing dot." example:"googleads.g.doubleclick.net."`
	Count  int64  `json:"count" doc:"Lookups over the window." example:"18"`
}

type listClientDomainsInput struct {
	dto.WindowInput
	Client string `path:"mac" doc:"MAC address in lowercase, with hyphens." example:"bc-07-1d-35-46-1f"`
	Status string `query:"status" enum:"allowed,blocked" default:"allowed" doc:"Whether the resolver allowed or blocked the lookups."`
	Limit  int    `query:"limit" minimum:"1" maximum:"100" default:"20" doc:"Number of domains."`
}

type listClientDomainsOutput struct {
	Body struct {
		Items []domainBody `json:"items" nullable:"false" doc:"The most looked-up domains first."`
	}
}

func registerListClientDomains(api huma.API, problems *problem.Mapper, clients *appnetwork.Clients) {
	huma.Register(api, huma.Operation{
		OperationID: "listClientDomains",
		Method:      http.MethodGet,
		Path:        "/clients/{mac}/domains",
		Summary:     "List the domains a client looked up",
		Tags:        []string{"network"},
	}, func(ctx context.Context, in *listClientDomainsInput) (*listClientDomainsOutput, error) {
		domains, err := clients.Domains(ctx, in.Client, domainnetwork.DomainStatus(in.Status), in.Limit, in.Start, in.End)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listClientDomainsOutput{}
		out.Body.Items = make([]domainBody, 0, len(domains))
		for _, d := range domains {
			out.Body.Items = append(out.Body.Items, domainBody{Domain: d.Domain, Count: d.Count})
		}
		return out, nil
	})
}
