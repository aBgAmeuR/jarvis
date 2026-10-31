package maintenance

type DriftKind string

const (
	MissingInIaC     DriftKind = "missingInIac"
	MissingInProxmox DriftKind = "missingInProxmox"
	ResourceMismatch DriftKind = "resourceMismatch"
)

type DriftItem struct {
	Kind  DriftKind
	Guest GuestRef
	// Differences is set for ResourceMismatch only.
	Differences []ResourceDifference
}
