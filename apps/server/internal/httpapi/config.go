package httpapi

import "github.com/danielgtaylor/huma/v2"

const basePath = "/api/v1"

func newConfig(version string) huma.Config {
	config := huma.DefaultConfig("Jarvis API", version)
	config.Info.Description = "Read-only console over the homelab observability stack."
	config.Servers = []*huma.Server{{URL: basePath}}
	config.DocsPath = ""
	return config
}
