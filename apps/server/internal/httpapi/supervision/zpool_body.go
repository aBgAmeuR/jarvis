package supervision

import (
	"time"

	domainsupervision "jarvis/server/internal/domain/supervision"
)

type zpoolBody struct {
	Name             string     `json:"name" doc:"Pool name." example:"hdd"`
	State            string     `json:"state" doc:"ZFS pool state." enum:"online,degraded,faulted,offline,removed,suspended,unavail,unknown" example:"online"`
	LatestSnapshotAt *time.Time `json:"latestSnapshotAt,omitempty" doc:"Creation time of the newest snapshot of the pool or of a dataset below it, RFC 3339 in UTC. Absent without any." example:"2026-10-31T00:00:00Z"`
}

func newZPoolBody(p domainsupervision.ZPool) zpoolBody {
	return zpoolBody{Name: p.Name, State: string(p.State), LatestSnapshotAt: p.LatestSnapshotAt}
}
