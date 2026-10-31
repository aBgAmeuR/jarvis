package supervision

import (
	"github.com/danielgtaylor/huma/v2"

	appsupervision "jarvis/server/internal/app/supervision"
	"jarvis/server/internal/httpapi/problem"
)

type Services struct {
	Guests   *appsupervision.Guests
	Nodes    *appsupervision.Nodes
	Metrics  *appsupervision.Metrics
	Hardware *appsupervision.Hardware
}

func Register(api huma.API, problems *problem.Mapper, s Services) {
	registerListGuests(api, problems, s.Guests)
	registerGetGuest(api, problems, s.Guests)
	registerGetGuestMetrics(api, problems, s.Metrics)
	registerListNodes(api, problems, s.Nodes)
	registerGetNode(api, problems, s.Nodes)
	registerGetNodeMetrics(api, problems, s.Metrics)
	registerListStorages(api, problems, s.Hardware)
	registerListZPools(api, problems, s.Hardware)
	registerListDisks(api, problems, s.Hardware)
}
