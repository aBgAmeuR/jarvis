package supervision

type Nodes struct {
	inventory InventorySource
}

func NewNodes(inventory InventorySource) *Nodes {
	return &Nodes{inventory: inventory}
}
