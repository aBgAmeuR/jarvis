package supervision

type Catalog struct {
	Guests []CatalogGuest
}

type CatalogGuest struct {
	VMID        int
	Address     string
	Description string
}
