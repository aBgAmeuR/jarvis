package bootstrap

import (
	"net/http"

	"jarvis/server/internal/platform/config"
	"jarvis/server/internal/platform/httpclient"
	"jarvis/server/internal/platform/telemetry"
)

const maxBodyBytes = 8 << 20

type clients struct {
	plain   *http.Client
	proxmox *http.Client
}

func newClients(cfg config.Config, metrics *telemetry.Metrics) (clients, error) {
	options := httpclient.Options{
		CacheTTL:     cfg.CacheTTL,
		Timeout:      cfg.RequestTimeout,
		MaxBodyBytes: maxBodyBytes,
		Wrap:         []func(http.RoundTripper) http.RoundTripper{metrics.MeasureSources, telemetry.TraceClient},
	}
	plain, err := httpclient.New(options)
	if err != nil {
		return clients{}, err
	}
	options.CAFile = cfg.ProxmoxCAFile
	proxmox, err := httpclient.New(options)
	if err != nil {
		return clients{}, err
	}
	return clients{plain: plain, proxmox: proxmox}, nil
}
