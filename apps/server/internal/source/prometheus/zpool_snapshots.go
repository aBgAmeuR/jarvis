package prometheus

import (
	"context"
	"strconv"
	"strings"
	"time"

	"jarvis/server/internal/source/failure"
)

func (c *Client) latestSnapshots(ctx context.Context, node string) (map[string]time.Time, error) {
	entries, err := c.queryVector(ctx, "zfs_snapshot_latest_timestamp_seconds{instance="+strconv.Quote(node)+"}")
	if err != nil {
		return nil, failure.Wrap(name, err)
	}
	latest := make(map[string]time.Time, len(entries))
	for _, e := range entries {
		latest[e.labels["dataset"]] = time.Unix(int64(e.value), 0).UTC()
	}
	return latest, nil
}

// A pool owns its own dataset and every dataset below it.
func poolSnapshot(pool string, datasets map[string]time.Time) *time.Time {
	var latest time.Time
	for dataset, at := range datasets {
		if (dataset == pool || strings.HasPrefix(dataset, pool+"/")) && at.After(latest) {
			latest = at
		}
	}
	if latest.IsZero() {
		return nil
	}
	return &latest
}
