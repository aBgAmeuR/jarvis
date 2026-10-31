package prometheus

import (
	"strings"

	"jarvis/server/internal/domain/availability"
)

// probeID is kind-vmid, kind-instance for the host, or http-site for HTTP probes.
func probeID(kind availability.ProbeKind, labels map[string]string) string {
	if kind == availability.ProbeHTTP {
		site, _, _ := strings.Cut(labels["site"], ".")
		return "http-" + site
	}
	if vmid := labels["vmid"]; vmid != "" {
		return string(kind) + "-" + vmid
	}
	return string(kind) + "-" + labels["instance"]
}
