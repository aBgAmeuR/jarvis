package bootstrap

import (
	"log/slog"

	apphealth "jarvis/server/internal/app/health"
	appidentity "jarvis/server/internal/app/identity"
	applogs "jarvis/server/internal/app/logs"
	"jarvis/server/internal/domain"
	"jarvis/server/internal/httpapi"
	"jarvis/server/internal/platform/config"
)

func newDeps(cfg config.Config, meta domain.Meta, logger *slog.Logger, src sources) httpapi.Deps {
	availability := newAvailability(src)
	services := newServiceFeature(src)
	network := newNetwork(src)
	return httpapi.Deps{
		Logger:       logger,
		Meta:         meta,
		Grafana:      domain.Grafana{URL: cfg.GrafanaURL, PrometheusUID: cfg.GrafanaPrometheusUID, LokiUID: cfg.GrafanaLokiUID},
		Identity:     appidentity.NewResolver(cfg.TrustedProxy),
		Health:       apphealth.NewService(cfg.RequestTimeout, src.checkers...),
		Search:       newSearch(src, services, availability.Probes, network.Clients, network.Peers),
		Overview:     newOverview(src, availability.Alerts, availability.Probes),
		Logs:         applogs.NewLogs(src.loki),
		Services:     services,
		Supervision:  newSupervision(src, availability.Alerts),
		Availability: availability,
		Network:      network,
		Maintenance:  newMaintenance(src),
	}
}
