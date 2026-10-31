package supervision

type Storage struct {
	Name       string
	Node       string
	Type       string
	Content    []string
	Status     string
	Shared     bool
	UsedBytes  uint64
	TotalBytes uint64
}
