package availability

import (
	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	"jarvis/server/internal/httpapi/problem"
)

type Services struct {
	Alerts       *appavailability.Alerts
	Probes       *appavailability.Probes
	Certificates *appavailability.Certificates
	Silences     *appavailability.Silences
}

func Register(api huma.API, problems *problem.Mapper, s Services) {
	registerListAlerts(api, problems, s.Alerts)
	registerGetAlert(api, problems, s.Alerts)
	registerListProbes(api, problems, s.Probes)
	registerGetProbe(api, problems, s.Probes)
	registerGetProbeMetrics(api, problems, s.Probes)
	registerGetProbeAvailability(api, problems, s.Probes)
	registerListCertificates(api, problems, s.Certificates)
	registerGetCertificate(api, problems, s.Certificates)
	registerListSilences(api, problems, s.Silences)
	registerGetSilence(api, problems, s.Silences)
}
