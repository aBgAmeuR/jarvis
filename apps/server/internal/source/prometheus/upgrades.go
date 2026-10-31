package prometheus

import (
	"context"
	"strconv"
	"sync"

	"jarvis/server/internal/domain/supervision"
	"jarvis/server/internal/source/failure"
)

const securityOrigin = `origin=~".*Debian-Security.*"`

// Origins listing the security repository count as security upgrades.
func (c *Client) Upgrades(ctx context.Context) (map[int]supervision.Upgrades, error) {
	exprs := [...]string{
		"sum by (vmid) (apt_upgrades_pending{" + securityOrigin + "})",
		"sum by (vmid) (apt_upgrades_pending{" + `origin!~".*Debian-Security.*"` + "})",
		"max by (vmid) (node_reboot_required)",
	}
	var results [len(exprs)][]entry
	var errs [len(exprs)]error
	var wg sync.WaitGroup
	for i, expr := range exprs {
		wg.Go(func() { results[i], errs[i] = c.queryVector(ctx, expr) })
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return nil, failure.Wrap(name, err)
		}
	}

	upgrades := map[int]supervision.Upgrades{}
	apply := func(entries []entry, set func(*supervision.Upgrades, float64)) {
		for _, e := range entries {
			vmid, err := strconv.Atoi(e.labels["vmid"])
			if err != nil {
				continue
			}
			u := upgrades[vmid]
			set(&u, e.value)
			upgrades[vmid] = u
		}
	}
	apply(results[0], func(u *supervision.Upgrades, v float64) { u.Security = int(v) })
	apply(results[1], func(u *supervision.Upgrades, v float64) { u.Regular = int(v) })
	apply(results[2], func(u *supervision.Upgrades, v float64) { u.RebootRequired = v == 1 })
	return upgrades, nil
}
