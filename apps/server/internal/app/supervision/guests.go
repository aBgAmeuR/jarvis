package supervision

type Guests struct {
	inventory InventorySource
	catalog   CatalogSource
	alerts    GuestAlertSource
	upgrades  UpgradeSource
}

func NewGuests(inventory InventorySource, catalog CatalogSource, alerts GuestAlertSource, upgrades UpgradeSource) *Guests {
	return &Guests{inventory: inventory, catalog: catalog, alerts: alerts, upgrades: upgrades}
}
