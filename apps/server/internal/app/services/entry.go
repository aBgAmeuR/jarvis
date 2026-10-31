package services

import (
	"strings"

	"jarvis/server/internal/domain/services"
)

// Slugs and dependencies come later, once every name is known.
func entries(catalog services.Catalog, signals services.Signals) []services.Service {
	active := map[unitKey]bool{}
	for _, u := range signals.Units {
		active[unitKey{u.VMID, u.Unit}] = u.Active
	}
	var out []services.Service
	for _, guest := range catalog.Guests {
		for _, unit := range guest.Units {
			status := services.StatusUnknown
			if up, known := active[unitKey{guest.VMID, unit}]; known {
				status = statusOf(up)
			}
			out = append(out, services.Service{
				Name: strings.TrimSuffix(unit, ".service"), Type: services.TypeSystemd,
				VMID: guest.VMID, GuestName: guest.Name, Status: status,
			})
		}
	}
	for _, c := range signals.Containers {
		out = append(out, services.Service{
			Name: c.Name, Type: services.TypeContainer, VMID: c.VMID,
			GuestName: guestName(catalog, c.VMID), Status: statusOf(c.OK),
		})
	}
	for _, site := range catalog.Sites {
		name, _, _ := strings.Cut(site.Host, ".")
		status := services.StatusUnknown
		if up, known := signals.Sites[site.Host]; known {
			status = statusOf(up)
		}
		out = append(out, services.Service{
			Name: name, Type: services.TypeSite, VMID: site.VMID, GuestName: guestName(catalog, site.VMID),
			Status: status, URL: site.URL, ProbeID: "http-" + name,
		})
	}
	return out
}

type unitKey struct {
	vmid int
	unit string
}

func statusOf(up bool) services.Status {
	if up {
		return services.StatusUp
	}
	return services.StatusDown
}

func guestName(catalog services.Catalog, vmid int) string {
	for _, g := range catalog.Guests {
		if g.VMID == vmid {
			return g.Name
		}
	}
	return ""
}
