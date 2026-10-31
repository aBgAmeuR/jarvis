package supervision

import "time"

type ZPoolState string

const (
	ZPoolOnline    ZPoolState = "online"
	ZPoolDegraded  ZPoolState = "degraded"
	ZPoolFaulted   ZPoolState = "faulted"
	ZPoolOffline   ZPoolState = "offline"
	ZPoolRemoved   ZPoolState = "removed"
	ZPoolSuspended ZPoolState = "suspended"
	ZPoolUnavail   ZPoolState = "unavail"
	ZPoolUnknown   ZPoolState = "unknown"
)

type ZPool struct {
	Name  string
	State ZPoolState
	// LatestSnapshotAt is nil when no dataset of the pool has a snapshot.
	LatestSnapshotAt *time.Time
}
