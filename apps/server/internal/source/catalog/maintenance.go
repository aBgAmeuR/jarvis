package catalog

import (
	"context"

	"jarvis/server/internal/domain/maintenance"
)

func (f *File) CatalogGuests(context.Context) ([]maintenance.GuestRef, error) {
	doc, err := f.read()
	if err != nil {
		return nil, err
	}
	guests := make([]maintenance.GuestRef, 0, len(doc.Guests))
	for _, g := range doc.Guests {
		ref := maintenance.GuestRef{VMID: g.VMID, Name: g.Name}
		if r := g.Resources; r != nil {
			ref.Resources = &maintenance.Resources{Cores: r.Cores, MemoryMB: r.MemoryMB, SwapMB: r.SwapMB, DiskGB: r.DiskGB}
		}
		guests = append(guests, ref)
	}
	return guests, nil
}
