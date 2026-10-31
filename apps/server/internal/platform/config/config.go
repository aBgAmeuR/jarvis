package config

import (
	"fmt"
	"log/slog"
	"net/netip"
	"net/url"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
)

type Secret string

func (Secret) LogValue() slog.Value { return slog.StringValue("***") }

func (Secret) String() string { return "***" }

func (Secret) MarshalText() ([]byte, error) { return []byte("***"), nil }

type Config struct {
	ListenAddr      string        `env:"JARVIS_LISTEN_ADDR" envDefault:":8080"`
	PrometheusURL   string        `env:"JARVIS_PROMETHEUS_URL" envDefault:"http://prometheus:9090"`
	LokiURL         string        `env:"JARVIS_LOKI_URL" envDefault:"http://loki:3100"`
	TempoURL        string        `env:"JARVIS_TEMPO_URL" envDefault:"http://tempo:3200"`
	AlertmanagerURL string        `env:"JARVIS_ALERTMANAGER_URL" envDefault:"http://alertmanager:9093"`
	ProxmoxURL      string        `env:"JARVIS_PROXMOX_URL,required"`
	ProxmoxToken    Secret        `env:"JARVIS_PROXMOX_TOKEN_FILE,required,file"`
	ProxmoxCAFile   string        `env:"JARVIS_PROXMOX_CA_FILE"`
	CatalogPath     string        `env:"JARVIS_CATALOG_PATH" envDefault:"/etc/jarvis/catalog.json"`
	TrustedProxy    netip.Addr    `env:"JARVIS_TRUSTED_PROXY,required"`
	FilesRoot       string        `env:"JARVIS_FILES_ROOT"`
	RequestTimeout  time.Duration `env:"JARVIS_REQUEST_TIMEOUT" envDefault:"5s"`
	CacheTTL        time.Duration `env:"JARVIS_CACHE_TTL" envDefault:"2s"`
	LogLevel        slog.Level    `env:"JARVIS_LOG_LEVEL" envDefault:"info"`
	OTLPEndpoint    string        `env:"JARVIS_OTLP_ENDPOINT"`
}

func Load() (Config, error) {
	var c Config
	if err := env.Parse(&c); err != nil {
		return Config{}, fmt.Errorf("read environment: %w", err)
	}
	c.ProxmoxToken = Secret(strings.TrimSpace(string(c.ProxmoxToken)))
	if err := c.validate(); err != nil {
		return Config{}, err
	}
	return c, nil
}

func (c Config) FilesEnabled() bool { return c.FilesRoot != "" }

func (c Config) validate() error {
	urls := []struct{ name, value string }{
		{"JARVIS_PROMETHEUS_URL", c.PrometheusURL},
		{"JARVIS_LOKI_URL", c.LokiURL},
		{"JARVIS_TEMPO_URL", c.TempoURL},
		{"JARVIS_ALERTMANAGER_URL", c.AlertmanagerURL},
		{"JARVIS_PROXMOX_URL", c.ProxmoxURL},
	}
	if c.OTLPEndpoint != "" {
		urls = append(urls, struct{ name, value string }{"JARVIS_OTLP_ENDPOINT", c.OTLPEndpoint})
	}
	for _, u := range urls {
		parsed, err := url.Parse(u.value)
		if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") || parsed.Host == "" {
			return fmt.Errorf("%s: %q is not an http(s) URL", u.name, u.value)
		}
	}
	if c.ProxmoxToken == "" {
		return fmt.Errorf("JARVIS_PROXMOX_TOKEN_FILE: file is empty")
	}
	if !c.TrustedProxy.IsValid() {
		return fmt.Errorf("JARVIS_TRUSTED_PROXY: not an IP address")
	}
	if c.RequestTimeout <= 0 {
		return fmt.Errorf("JARVIS_REQUEST_TIMEOUT: must be positive")
	}
	if c.CacheTTL < 0 {
		return fmt.Errorf("JARVIS_CACHE_TTL: must not be negative")
	}
	return nil
}
