package prometheus

import (
	"context"
	"slices"
	"strconv"
	"strings"

	"jarvis/server/internal/domain/supervision"
	"jarvis/server/internal/source/failure"
)

// node_exporter exposes the state of a pool as one 0/1 series per state.
func (c *Client) ZPools(ctx context.Context, node string) ([]supervision.ZPool, error) {
	entries, err := c.queryVector(ctx, "node_zfs_zpool_state{instance="+strconv.Quote(node)+"}")
	if err != nil {
		return nil, failure.Wrap(name, err)
	}
	states := map[string]supervision.ZPoolState{}
	for _, e := range entries {
		pool := e.labels["zpool"]
		if _, seen := states[pool]; !seen {
			states[pool] = supervision.ZPoolUnknown
		}
		if e.value == 1 {
			states[pool] = supervision.ZPoolState(e.labels["state"])
		}
	}
	snapshots, err := c.latestSnapshots(ctx, node)
	if err != nil {
		return nil, err
	}
	pools := make([]supervision.ZPool, 0, len(states))
	for pool, state := range states {
		pools = append(pools, supervision.ZPool{Name: pool, State: state, LatestSnapshotAt: poolSnapshot(pool, snapshots)})
	}
	slices.SortFunc(pools, func(a, b supervision.ZPool) int { return strings.Compare(a.Name, b.Name) })
	return pools, nil
}
