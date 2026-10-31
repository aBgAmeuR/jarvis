package overview

import (
	"time"

	domainoverview "jarvis/server/internal/domain/overview"
	"jarvis/server/internal/httpapi/dto"
)

type overviewAlertBody struct {
	Fingerprint string    `json:"fingerprint" doc:"Alertmanager fingerprint, the id to read the alert with." example:"3a1f9c0d5b7e2a44"`
	Name        string    `json:"name" doc:"Alert rule name." example:"GuestDown"`
	Severity    string    `json:"severity" doc:"Severity label of the rule." enum:"critical,warning,info" example:"critical"`
	StartsAt    time.Time `json:"startsAt" doc:"When the alert started firing, RFC 3339 in UTC." example:"2026-10-31T09:41:00Z"`
	Summary     string    `json:"summary" doc:"Summary annotation of the alert. Empty without one." example:"Guest 116 (dns) is not running"`
}

type overviewUsageBody struct {
	CPUUsageRatio    float64 `json:"cpuUsageRatio" doc:"CPU in use, from 0 to 1 of all cores." example:"0.12"`
	MemoryUsedBytes  uint64  `json:"memoryUsedBytes" doc:"Memory in use." example:"1073741824"`
	MemoryTotalBytes uint64  `json:"memoryTotalBytes" doc:"Memory available." example:"4294967296"`
	DiskUsedBytes    uint64  `json:"diskUsedBytes" doc:"Root disk in use." example:"2147483648"`
	DiskTotalBytes   uint64  `json:"diskTotalBytes" doc:"Root disk size." example:"8589934592"`
}

type overviewNodeBody struct {
	Name   string             `json:"name" doc:"Proxmox node name." example:"pve"`
	Online bool               `json:"online" doc:"Whether Proxmox reports the node online." example:"true"`
	Usage  *overviewUsageBody `json:"usage,omitempty" doc:"Resource usage. Absent when the node is offline or Proxmox gives no figures."`
}

type overviewGuestsBody struct {
	Total   int `json:"total" doc:"Guests in Proxmox, templates excluded." example:"7"`
	Running int `json:"running" doc:"Guests that are running." example:"7"`
}

type overviewProbesBody struct {
	Total int `json:"total" doc:"Probes." example:"12"`
	Up    int `json:"up" doc:"Probes whose last check succeeded." example:"12"`
}

type overviewBody struct {
	Status   string              `json:"status" doc:"Severity of the worst firing alert, ok without any, unknown when Alertmanager did not answer." enum:"ok,info,warning,critical,unknown" example:"ok"`
	ToHandle []overviewAlertBody `json:"toHandle" nullable:"false" doc:"The firing alerts, the most severe first."`
	Nodes    []overviewNodeBody  `json:"nodes" nullable:"false" doc:"Proxmox nodes."`
	Guests   *overviewGuestsBody `json:"guests,omitempty" doc:"Guest counters. Absent when Proxmox did not answer."`
	Probes   *overviewProbesBody `json:"probes,omitempty" doc:"Probe counters. Absent when Prometheus did not answer."`
	Degraded []dto.DegradedBody  `json:"degraded" nullable:"false" doc:"Sources that failed. Their part is missing."`
}

func newOverviewBody(o domainoverview.Overview, degraded []dto.DegradedBody) overviewBody {
	body := overviewBody{
		Status:   string(o.Status),
		ToHandle: make([]overviewAlertBody, 0, len(o.ToHandle)),
		Nodes:    make([]overviewNodeBody, 0, len(o.Nodes)),
		Degraded: degraded,
	}
	for _, a := range o.ToHandle {
		body.ToHandle = append(body.ToHandle, overviewAlertBody{
			Fingerprint: a.Fingerprint,
			Name:        a.Name,
			Severity:    string(a.Severity),
			StartsAt:    a.StartsAt,
			Summary:     a.Summary,
		})
	}
	for _, n := range o.Nodes {
		body.Nodes = append(body.Nodes, newOverviewNodeBody(n))
	}
	if o.Guests != nil {
		body.Guests = &overviewGuestsBody{Total: o.Guests.Total, Running: o.Guests.Running}
	}
	if o.Probes != nil {
		body.Probes = &overviewProbesBody{Total: o.Probes.Total, Up: o.Probes.Up}
	}
	return body
}

func newOverviewNodeBody(n domainoverview.Node) overviewNodeBody {
	body := overviewNodeBody{Name: n.Name, Online: n.Online}
	if u := n.Usage; u != nil {
		body.Usage = &overviewUsageBody{
			CPUUsageRatio:    u.CPUUsageRatio,
			MemoryUsedBytes:  u.MemoryUsedBytes,
			MemoryTotalBytes: u.MemoryTotalBytes,
			DiskUsedBytes:    u.DiskUsedBytes,
			DiskTotalBytes:   u.DiskTotalBytes,
		}
	}
	return body
}
