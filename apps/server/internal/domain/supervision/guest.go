package supervision

import "jarvis/server/internal/domain"

type GuestKind string

const (
	GuestLXC  GuestKind = "lxc"
	GuestQEMU GuestKind = "qemu"
)

type Guest struct {
	VMID   int
	Name   string
	Kind   GuestKind
	Node   string
	Status string
	// Usage is nil unless the guest is running.
	Usage *Usage

	// From the catalog. ManagedByIaC is nil when the catalog is unavailable.
	Address      string
	Description  string
	ManagedByIaC *bool

	// From the firing alerts of the guest.
	AlertSeverity domain.Severity
	AlertCount    int

	// Upgrades is nil when the guest reports no apt figures.
	Upgrades *Upgrades
}
