package overview

import (
	"github.com/danielgtaylor/huma/v2"

	appoverview "jarvis/server/internal/app/overview"
)

func Register(api huma.API, overview *appoverview.Service) {
	registerGetOverview(api, overview)
}
