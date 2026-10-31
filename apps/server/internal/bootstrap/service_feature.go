package bootstrap

import appservices "jarvis/server/internal/app/services"

func newServiceFeature(s sources) *appservices.Services {
	return appservices.NewServices(s.catalog, s.prometheus)
}
