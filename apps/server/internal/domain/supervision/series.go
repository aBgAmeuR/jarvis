package supervision

// Series names are part of the API: the browser never sends PromQL.
type GuestSeries string

const (
	GuestCPUUsageRatio         GuestSeries = "cpuUsageRatio"
	GuestMemoryUsedBytes       GuestSeries = "memoryUsedBytes"
	GuestDiskUsedBytes         GuestSeries = "diskUsedBytes"
	GuestDiskReadBytesPerSec   GuestSeries = "diskReadBytesPerSecond"
	GuestDiskWriteBytesPerSec  GuestSeries = "diskWriteBytesPerSecond"
	GuestNetworkReceivePerSec  GuestSeries = "networkReceiveBytesPerSecond"
	GuestNetworkTransmitPerSec GuestSeries = "networkTransmitBytesPerSecond"
)

func AllGuestSeries() []GuestSeries {
	return []GuestSeries{
		GuestCPUUsageRatio, GuestMemoryUsedBytes, GuestDiskUsedBytes,
		GuestDiskReadBytesPerSec, GuestDiskWriteBytesPerSec,
		GuestNetworkReceivePerSec, GuestNetworkTransmitPerSec,
	}
}

type NodeSeries string

const (
	NodeCPUUsageRatio   NodeSeries = "cpuUsageRatio"
	NodeMemoryUsedBytes NodeSeries = "memoryUsedBytes"
	NodeDiskUsedBytes   NodeSeries = "diskUsedBytes"
	NodeLoad1           NodeSeries = "load1"
	NodeCPUTemperature  NodeSeries = "cpuTemperatureCelsius"
)

func AllNodeSeries() []NodeSeries {
	return []NodeSeries{NodeCPUUsageRatio, NodeMemoryUsedBytes, NodeDiskUsedBytes, NodeLoad1, NodeCPUTemperature}
}
