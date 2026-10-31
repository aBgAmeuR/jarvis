package availability

import domainavailability "jarvis/server/internal/domain/availability"

type probeBody struct {
	ID              string   `json:"id" doc:"Probe id: kind and guest, or http and site." example:"icmp-116"`
	Kind            string   `json:"kind" doc:"What the probe does." enum:"icmp,dns,http" example:"icmp"`
	Name            string   `json:"name" doc:"Name of the probed guest or host." example:"dns"`
	Target          string   `json:"target" doc:"Address or URL probed." example:"192.168.1.116"`
	VMID            *int     `json:"vmid,omitempty" doc:"Probed guest. Absent for the host." example:"116"`
	Up              bool     `json:"up" doc:"Whether the last probe succeeded." example:"true"`
	DurationSeconds float64  `json:"durationSeconds" doc:"How long the last probe took." example:"0.0012"`
	HTTPStatusCode  *int     `json:"httpStatusCode,omitempty" doc:"Status of the last HTTP probe. Absent for other kinds." example:"200"`
	Availability24h *float64 `json:"availability24h,omitempty" doc:"Share of successful probes over 24 hours, from 0 to 1. Absent without samples." example:"1"`
	Availability7d  *float64 `json:"availability7d,omitempty" doc:"Share of successful probes over 7 days, from 0 to 1. Absent without samples." example:"0.999"`
	Availability30d *float64 `json:"availability30d,omitempty" doc:"Share of successful probes over 30 days, from 0 to 1. Absent without samples." example:"0.998"`
}

func newProbeBody(p domainavailability.Probe) probeBody {
	return probeBody{
		ID:              p.ID,
		Kind:            string(p.Kind),
		Name:            p.Name,
		Target:          p.Target,
		VMID:            p.VMID,
		Up:              p.Up,
		DurationSeconds: p.Duration,
		HTTPStatusCode:  p.HTTPStatus,
		Availability24h: p.Availability24h,
		Availability7d:  p.Availability7d,
		Availability30d: p.Availability30d,
	}
}
