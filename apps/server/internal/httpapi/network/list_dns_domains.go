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

type listDNSDomainsInput struct {
	dto.WindowInput
	Status string `query:"status" enum:"allowed,blocked" default:"allowed" doc:"Whether the resolver allowed or blocked the lookups."`
	Limit  int    `query:"limit" minimum:"1" maximum:"100" default:"20" doc:"Number of domains."`
}

type listDNSDomainsOutput struct {
	Body struct {
		Items []domainBody `json:"items" nullable:"false" doc:"The most looked-up domains first."`
	}
}

func registerListDNSDomains(api huma.API, problems *problem.Mapper, dns *appnetwork.DNS) {
	huma.Register(api, huma.Operation{
		OperationID: "listDnsDomains",
		Method:      http.MethodGet,
		Path:        "/dns/domains",
		Summary:     "List the domains the resolver was asked for",
		Tags:        []string{"network"},
	}, func(ctx context.Context, in *listDNSDomainsInput) (*listDNSDomainsOutput, error) {
		domains, err := dns.Domains(ctx, domainnetwork.DomainStatus(in.Status), in.Limit, in.Start, in.End)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listDNSDomainsOutput{}
		out.Body.Items = make([]domainBody, 0, len(domains))
		for _, d := range domains {
			out.Body.Items = append(out.Body.Items, domainBody{Domain: d.Domain, Count: d.Count})
		}
		return out, nil
	})
}
