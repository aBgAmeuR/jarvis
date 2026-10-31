package services

type SystemdUnit struct {
	VMID   int
	Unit   string
	Active bool
}

type Container struct {
	VMID int
	Name string
	OK   bool
}

// Edge says that Client calls Server, as the service graph names them.
type Edge struct {
	Client string
	Server string
}

type Signals struct {
	Units      []SystemdUnit
	Containers []Container
	// Sites maps a site host to whether its last probe succeeded.
	Sites        map[string]bool
	Edges        []Edge
	Instrumented map[string]bool
}
