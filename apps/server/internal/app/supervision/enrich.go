package supervision

import (
	"slices"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/supervision"
)

// A nil catalog leaves ManagedByIaC unset: a missing entry would then mean nothing.
func enrich(guests []supervision.Guest, catalog *supervision.Catalog, alerts []supervision.GuestAlert, upgrades map[int]supervision.Upgrades) []supervision.Guest {
	entries := map[int]supervision.CatalogGuest{}
	if catalog != nil {
		for _, entry := range catalog.Guests {
			entries[entry.VMID] = entry
		}
	}
	worst := map[int]domain.Severity{}
	count := map[int]int{}
	for _, alert := range alerts {
		if current, seen := worst[alert.VMID]; !seen || alert.Severity.Rank() < current.Rank() {
			worst[alert.VMID] = alert.Severity
		}
		count[alert.VMID]++
	}

	out := slices.Clone(guests)
	for i := range out {
		guest := &out[i]
		if catalog != nil {
			entry, managed := entries[guest.VMID]
			guest.Address, guest.Description = entry.Address, entry.Description
			guest.ManagedByIaC = &managed
		}
		guest.AlertSeverity, guest.AlertCount = worst[guest.VMID], count[guest.VMID]
		if u, ok := upgrades[guest.VMID]; ok {
			guest.Upgrades = &u
		}
	}
	slices.SortFunc(out, func(x, y supervision.Guest) int { return x.VMID - y.VMID })
	return out
}
