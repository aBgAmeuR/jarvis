package supervision

type Usage struct {
	CPUUsageRatio    float64
	CPUCount         int
	MemoryUsedBytes  uint64
	MemoryTotalBytes uint64
	DiskUsedBytes    uint64
	DiskTotalBytes   uint64
	UptimeSeconds    uint64
}
