package overview

type Usage struct {
	CPUUsageRatio    float64
	MemoryUsedBytes  uint64
	MemoryTotalBytes uint64
	DiskUsedBytes    uint64
	DiskTotalBytes   uint64
}

type Node struct {
	Name   string
	Online bool
	// Usage is nil when the node is offline or the source gives no figures.
	Usage *Usage
}

type Guests struct {
	Total   int
	Running int
}

type Inventory struct {
	Nodes  []Node
	Guests Guests
}
