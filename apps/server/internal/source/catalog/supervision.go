package catalog

import (
	"context"

	"jarvis/server/internal/domain/supervision"
)

func (f *File) Catalog(context.Context) (supervision.Catalog, error) {
	doc, err := f.read()
	if err != nil {
		return supervision.Catalog{}, err
	}
	catalog := supervision.Catalog{Guests: make([]supervision.CatalogGuest, 0, len(doc.Guests))}
	for _, g := range doc.Guests {
		catalog.Guests = append(catalog.Guests, supervision.CatalogGuest{VMID: g.VMID, Address: g.Address, Description: g.Description})
	}
	return catalog, nil
}
