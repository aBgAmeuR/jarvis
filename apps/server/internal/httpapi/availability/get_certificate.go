package availability

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	"jarvis/server/internal/httpapi/problem"
)

type getCertificateInput struct {
	Site string `path:"site" doc:"Site host name." example:"files.antoinejosset.fr"`
}

type getCertificateOutput struct {
	Body certificateBody
}

func registerGetCertificate(api huma.API, problems *problem.Mapper, certificates *appavailability.Certificates) {
	huma.Register(api, huma.Operation{
		OperationID: "getCertificate",
		Method:      http.MethodGet,
		Path:        "/certificates/{site}",
		Summary:     "Get a certificate",
		Tags:        []string{"availability"},
	}, func(ctx context.Context, in *getCertificateInput) (*getCertificateOutput, error) {
		certificate, err := certificates.Get(ctx, in.Site)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		return &getCertificateOutput{Body: newCertificateBody(certificate)}, nil
	})
}
