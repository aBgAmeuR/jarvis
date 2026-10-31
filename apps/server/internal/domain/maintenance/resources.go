package maintenance

type ResourceField string

const (
	FieldCores    ResourceField = "cores"
	FieldMemoryMB ResourceField = "memoryMb"
	FieldSwapMB   ResourceField = "swapMb"
	FieldDiskGB   ResourceField = "diskGb"
)

type Resources struct {
	Cores    int
	MemoryMB int
	SwapMB   int
	DiskGB   int
}

type ResourceDifference struct {
	Field    ResourceField
	Declared int
	Actual   int
}

// Differences compares r, the declared resources, with actual.
func (r Resources) Differences(actual Resources) []ResourceDifference {
	var diffs []ResourceDifference
	for _, f := range []struct {
		field            ResourceField
		declared, actual int
	}{
		{FieldCores, r.Cores, actual.Cores},
		{FieldMemoryMB, r.MemoryMB, actual.MemoryMB},
		{FieldSwapMB, r.SwapMB, actual.SwapMB},
		{FieldDiskGB, r.DiskGB, actual.DiskGB},
	} {
		if f.declared != f.actual {
			diffs = append(diffs, ResourceDifference{Field: f.field, Declared: f.declared, Actual: f.actual})
		}
	}
	return diffs
}
