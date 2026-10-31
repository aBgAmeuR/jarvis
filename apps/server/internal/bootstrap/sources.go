package bootstrap

import (
	"jarvis/server/internal/app/health"
	"jarvis/server/internal/platform/config"
	"jarvis/server/internal/platform/httpclient"
	"jarvis/server/internal/source/alertmanager"
	"jarvis/server/internal/source/catalog"
	"jarvis/server/internal/source/loki"
	"jarvis/server/internal/source/pathcheck"
	"jarvis/server/internal/source/prometheus"
	"jarvis/server/internal/source/proxmox"
	"jarvis/server/internal/source/readiness"
)

type sources struct {
	alertmanager *alertmanager.Client
	prometheus   *prometheus.Client
	loki         *loki.Client
	proxmox      *proxmox.Client
	catalog      *catalog.File
	checkers     []health.Checker
}

func newSources(cfg config.Config, c clients) sources {
	fetcher := func(baseURL string) *httpclient.Fetcher {
		return httpclient.NewFetcher(c.plain, baseURL, "", cfg.RequestTimeout, maxBodyBytes)
	}
	s := sources{
		alertmanager: alertmanager.New(fetcher(cfg.AlertmanagerURL)),
		prometheus:   prometheus.New(fetcher(cfg.PrometheusURL)),
		loki:         loki.New(fetcher(cfg.LokiURL)),
		proxmox: proxmox.New(httpclient.NewFetcher(
			c.proxmox, cfg.ProxmoxURL, "PVEAPIToken="+string(cfg.ProxmoxToken), cfg.RequestTimeout, maxBodyBytes,
		)),
		catalog: catalog.NewFile(cfg.CatalogPath),
	}
	s.checkers = []health.Checker{
		s.prometheus,
		s.loki,
		readiness.New("tempo", "/ready", fetcher(cfg.TempoURL)),
		s.alertmanager,
		s.proxmox,
		s.catalog,
	}
	if cfg.FilesEnabled() {
		s.checkers = append(s.checkers, pathcheck.New("files", cfg.FilesRoot))
	}
	return s
}
