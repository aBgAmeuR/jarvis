package catalog

import (
	"context"

	"jarvis/server/internal/domain/services"
)

func (f *File) ServiceCatalog(context.Context) (services.Catalog, error) {
	doc, err := f.read()
	if err != nil {
		return services.Catalog{}, err
	}
	catalog := services.Catalog{
		Guests: make([]services.CatalogGuest, 0, len(doc.Guests)),
		Sites:  make([]services.CatalogSite, 0, len(doc.Sites)),
	}
	for _, g := range doc.Guests {
		catalog.Guests = append(catalog.Guests, services.CatalogGuest{VMID: g.VMID, Name: g.Name, Units: g.Units})
	}
	for _, s := range doc.Sites {
		site := services.CatalogSite{Host: s.Host, URL: s.URL}
		if s.VMID != nil {
			site.VMID = *s.VMID
		}
		catalog.Sites = append(catalog.Sites, site)
	}
	return catalog, nil
}
