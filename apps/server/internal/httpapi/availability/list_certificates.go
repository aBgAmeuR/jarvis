package availability

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	"jarvis/server/internal/httpapi/problem"
)

type listCertificatesOutput struct {
	Body struct {
		Items []certificateBody `json:"items" nullable:"false" doc:"Certificates, the first to expire first."`
	}
}

func registerListCertificates(api huma.API, problems *problem.Mapper, certificates *appavailability.Certificates) {
	huma.Register(api, huma.Operation{
		OperationID: "listCertificates",
		Method:      http.MethodGet,
		Path:        "/certificates",
		Summary:     "List certificates",
		Tags:        []string{"availability"},
	}, func(ctx context.Context, _ *struct{}) (*listCertificatesOutput, error) {
		list, err := certificates.List(ctx)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listCertificatesOutput{}
		out.Body.Items = make([]certificateBody, 0, len(list))
		for _, c := range list {
			out.Body.Items = append(out.Body.Items, newCertificateBody(c))
		}
		return out, nil
	})
}
