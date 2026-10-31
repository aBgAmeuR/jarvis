package availability

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	domainavailability "jarvis/server/internal/domain/availability"
	"jarvis/server/internal/httpapi/problem"
)

type listSilencesInput struct {
	State string `query:"state" enum:"active,pending,expired" doc:"Only silences in this state. All when absent."`
}

type listSilencesOutput struct {
	Body struct {
		Items []silenceBody `json:"items" nullable:"false" doc:"Silences, the last to end first."`
	}
}

func registerListSilences(api huma.API, problems *problem.Mapper, silences *appavailability.Silences) {
	huma.Register(api, huma.Operation{
		OperationID: "listSilences",
		Method:      http.MethodGet,
		Path:        "/silences",
		Summary:     "List silences",
		Tags:        []string{"availability"},
	}, func(ctx context.Context, in *listSilencesInput) (*listSilencesOutput, error) {
		list, err := silences.List(ctx, domainavailability.SilenceState(in.State))
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &listSilencesOutput{}
		out.Body.Items = make([]silenceBody, 0, len(list))
		for _, s := range list {
			out.Body.Items = append(out.Body.Items, newSilenceBody(s))
		}
		return out, nil
	})
}
