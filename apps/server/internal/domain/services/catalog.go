package services

type Catalog struct {
	Guests []CatalogGuest
	Sites  []CatalogSite
}

type CatalogGuest struct {
	VMID  int
	Name  string
	Units []string
}

type CatalogSite struct {
	Host string
	URL  string
	VMID int
}
