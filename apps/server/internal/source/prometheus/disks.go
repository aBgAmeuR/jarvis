package prometheus

import (
	"context"
	"slices"
	"strconv"
	"strings"

	"jarvis/server/internal/domain/supervision"
	"jarvis/server/internal/source/failure"
)

func (c *Client) Disks(ctx context.Context, node string) ([]supervision.Disk, error) {
	entries, err := c.queryVector(ctx, `{__name__=~"smart_.+",instance=`+strconv.Quote(node)+`}`)
	if err != nil {
		return nil, failure.Wrap(name, err)
	}
	info := map[string]map[string]string{}
	values := map[string]map[string]float64{}
	for _, e := range entries {
		device, metric := e.labels["device"], e.labels["__name__"]
		if metric == "smart_device_info" {
			info[device] = e.labels
			continue
		}
		if values[device] == nil {
			values[device] = map[string]float64{}
		}
		values[device][metric] = e.value
	}

	disks := make([]supervision.Disk, 0, len(info))
	for device, labels := range info {
		disks = append(disks, newDisk(device, labels, values[device]))
	}
	slices.SortFunc(disks, func(a, b supervision.Disk) int { return strings.Compare(a.Device, b.Device) })
	return disks, nil
}
