package logs

import (
	"github.com/danielgtaylor/huma/v2"

	applogs "jarvis/server/internal/app/logs"
	"jarvis/server/internal/httpapi/problem"
)

func Register(api huma.API, problems *problem.Mapper, logs *applogs.Logs) {
	registerSearchLogs(api, problems, logs)
	registerVolumeLogs(api, problems, logs)
}
