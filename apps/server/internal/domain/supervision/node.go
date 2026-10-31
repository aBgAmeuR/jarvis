package supervision

type Node struct {
	Name   string
	Online bool
	// Usage is nil when the node is offline or the source gives no figures.
	Usage *Usage
}
