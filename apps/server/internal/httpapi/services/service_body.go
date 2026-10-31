package services

import domainservices "jarvis/server/internal/domain/services"

type ServiceBody struct {
	Slug         string   `json:"slug" doc:"Service id: its name, with the guest added when several share it." example:"grafana"`
	Name         string   `json:"name" doc:"Unit, container or site name." example:"grafana"`
	Type         string   `json:"type" doc:"Where the service comes from." enum:"systemd,container,site" example:"container"`
	VMID         int      `json:"vmid" doc:"Guest that hosts it. Zero for a site without a guest." example:"111"`
	GuestName    string   `json:"guestName,omitempty" doc:"Name of that guest." example:"monitoring"`
	Status       string   `json:"status" doc:"Whether monitoring sees the service running." enum:"up,down,unknown" example:"up"`
	URL          string   `json:"url,omitempty" doc:"Public URL. Sites only." example:"https://monitoring.antoinejosset.fr"`
	ProbeID      string   `json:"probeId,omitempty" doc:"Probe that checks the site. Sites only." example:"http-monitoring"`
	DependsOn    []string `json:"dependsOn" nullable:"false" doc:"Slugs of the services it calls, from the service graph." example:"[\"grafana\"]"`
	Instrumented bool     `json:"instrumented" doc:"Whether request metrics exist for it." example:"true"`
}

func newServiceBody(s domainservices.Service) ServiceBody {
	dependsOn := s.DependsOn
	if dependsOn == nil {
		dependsOn = []string{}
	}
	return ServiceBody{
		Slug:         s.Slug,
		Name:         s.Name,
		Type:         string(s.Type),
		VMID:         s.VMID,
		GuestName:    s.GuestName,
		Status:       string(s.Status),
		URL:          s.URL,
		ProbeID:      s.ProbeID,
		DependsOn:    dependsOn,
		Instrumented: s.Instrumented,
	}
}
