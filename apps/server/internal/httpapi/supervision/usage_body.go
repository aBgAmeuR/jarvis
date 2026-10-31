package supervision

import domainsupervision "jarvis/server/internal/domain/supervision"

type usageBody struct {
	CPUUsageRatio    float64 `json:"cpuUsageRatio" doc:"CPU in use, from 0 to 1 of all cores." example:"0.12"`
	CPUCount         int     `json:"cpuCount" doc:"Number of cores." example:"2"`
	MemoryUsedBytes  uint64  `json:"memoryUsedBytes" doc:"Memory in use." example:"1073741824"`
	MemoryTotalBytes uint64  `json:"memoryTotalBytes" doc:"Memory available." example:"4294967296"`
	DiskUsedBytes    uint64  `json:"diskUsedBytes" doc:"Root disk in use." example:"2147483648"`
	DiskTotalBytes   uint64  `json:"diskTotalBytes" doc:"Root disk size." example:"8589934592"`
	UptimeSeconds    uint64  `json:"uptimeSeconds" doc:"Time since boot." example:"86400"`
}

func newUsageBody(u *domainsupervision.Usage) *usageBody {
	if u == nil {
		return nil
	}
	return &usageBody{
		CPUUsageRatio:    u.CPUUsageRatio,
		CPUCount:         u.CPUCount,
		MemoryUsedBytes:  u.MemoryUsedBytes,
		MemoryTotalBytes: u.MemoryTotalBytes,
		DiskUsedBytes:    u.DiskUsedBytes,
		DiskTotalBytes:   u.DiskTotalBytes,
		UptimeSeconds:    u.UptimeSeconds,
	}
}
