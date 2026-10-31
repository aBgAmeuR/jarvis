package services

type Type string

const (
	TypeSystemd   Type = "systemd"
	TypeContainer Type = "container"
	TypeSite      Type = "site"
)

type Status string

const (
	StatusUp      Status = "up"
	StatusDown    Status = "down"
	StatusUnknown Status = "unknown"
)

type Service struct {
	Slug      string
	Name      string
	Type      Type
	VMID      int
	GuestName string
	Status    Status
	// URL and ProbeID are set for sites only.
	URL          string
	ProbeID      string
	DependsOn    []string
	Instrumented bool
}
