package maintenance

type GuestRef struct {
	VMID int
	Name string
	Kind string
	// Resources are declared by the IaC. Nil from Proxmox, or when the catalog does not carry them.
	Resources *Resources
}
