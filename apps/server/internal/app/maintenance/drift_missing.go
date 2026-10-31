package maintenance

import (
	"slices"

	"jarvis/server/internal/domain/maintenance"
)

func hasGuest(list []maintenance.GuestRef, vmid int) bool {
	return slices.ContainsFunc(list, func(g maintenance.GuestRef) bool { return g.VMID == vmid })
}

func missingItems(inProxmox, inCatalog []maintenance.GuestRef) []maintenance.DriftItem {
	var items []maintenance.DriftItem
	for _, g := range inProxmox {
		if !hasGuest(inCatalog, g.VMID) {
			items = append(items, maintenance.DriftItem{Kind: maintenance.MissingInIaC, Guest: g})
		}
	}
	for _, g := range inCatalog {
		if !hasGuest(inProxmox, g.VMID) {
			items = append(items, maintenance.DriftItem{Kind: maintenance.MissingInProxmox, Guest: g})
		}
	}
	return items
}
