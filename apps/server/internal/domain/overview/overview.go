package overview

// A nil part comes from a source that failed.
type Overview struct {
	Status Status
	// ToHandle is the firing alerts, the most severe first.
	ToHandle []Alert
	Nodes    []Node
	Guests   *Guests
	Probes   *Probes
}
