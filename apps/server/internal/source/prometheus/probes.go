package prometheus

import (
	"context"
	"slices"
	"strconv"
	"strings"
	"sync"

	"jarvis/server/internal/domain/availability"
	"jarvis/server/internal/source/failure"
)

const probeSelector = `{job=~"probe_.+"}`

func probeKey(labels map[string]string) string { return labels["job"] + "|" + labels["target"] }

// Probes joins the blackbox series on their job and target labels.
func (c *Client) Probes(ctx context.Context) ([]availability.Probe, error) {
	exprs := [...]string{
		"probe_success" + probeSelector,
		"probe_duration_seconds" + probeSelector,
		"probe_http_status_code" + probeSelector,
		"avg_over_time(probe_success" + probeSelector + "[24h])",
		"avg_over_time(probe_success" + probeSelector + "[7d])",
		"avg_over_time(probe_success" + probeSelector + "[30d])",
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

	byKey := func(entries []entry) map[string]float64 {
		out := make(map[string]float64, len(entries))
		for _, e := range entries {
			out[probeKey(e.labels)] = e.value
		}
		return out
	}
	duration, status := byKey(results[1]), byKey(results[2])
	day, week, month := byKey(results[3]), byKey(results[4]), byKey(results[5])
	ratio := func(values map[string]float64, key string) *float64 {
		if v, ok := values[key]; ok {
			return &v
		}
		return nil
	}

	probes := make([]availability.Probe, 0, len(results[0]))
	for _, e := range results[0] {
		key := probeKey(e.labels)
		kind := availability.ProbeKind(strings.TrimPrefix(e.labels["job"], "probe_"))
		probe := availability.Probe{
			ID:              probeID(kind, e.labels),
			Kind:            kind,
			Name:            e.labels["instance"],
			Target:          e.labels["target"],
			Up:              e.value == 1,
			Duration:        duration[key],
			Availability24h: ratio(day, key),
			Availability7d:  ratio(week, key),
			Availability30d: ratio(month, key),
		}
		if vmid, err := strconv.Atoi(e.labels["vmid"]); err == nil {
			probe.VMID = &vmid
		}
		if code, ok := status[key]; ok {
			n := int(code)
			probe.HTTPStatus = &n
		}
		probes = append(probes, probe)
	}
	slices.SortFunc(probes, func(a, b availability.Probe) int { return strings.Compare(a.ID, b.ID) })
	return probes, nil
}
