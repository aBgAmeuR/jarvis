package identity

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"jarvis/server/internal/httpapi/middleware"
)

type MeBody struct {
	User    string   `json:"user" doc:"Account name forwarded by TinyAuth." example:"antoine"`
	Groups  []string `json:"groups" nullable:"false" doc:"Groups of the account." example:"[\"admin\"]"`
	IsAdmin bool     `json:"isAdmin" doc:"Whether the account is in the admin group." example:"true"`
}

type getMeOutput struct {
	Body MeBody
}

func Register(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "getMe",
		Method:      http.MethodGet,
		Path:        "/me",
		Summary:     "Get the logged-in account",
		Tags:        []string{"common"},
	}, func(ctx context.Context, _ *struct{}) (*getMeOutput, error) {
		user, ok := middleware.UserFrom(ctx)
		if !ok {
			return nil, huma.Error401Unauthorized("Unauthorized")
		}
		return &getMeOutput{Body: MeBody{User: user.Name, Groups: user.Groups, IsAdmin: user.IsAdmin()}}, nil
	})
}
