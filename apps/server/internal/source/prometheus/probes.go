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

	series := probeSeries{
		duration: byProbeKey(results[1]),
		status:   byProbeKey(results[2]),
		day:      byProbeKey(results[3]),
		week:     byProbeKey(results[4]),
		month:    byProbeKey(results[5]),
	}
	probes := make([]availability.Probe, 0, len(results[0]))
	for _, e := range results[0] {
		probes = append(probes, newProbe(e, series))
	}
	slices.SortFunc(probes, func(a, b availability.Probe) int { return strings.Compare(a.ID, b.ID) })
	return probes, nil
}

// probeSeries holds the series joined to probe_success, keyed by probeKey.
type probeSeries struct {
	duration, status, day, week, month map[string]float64
}

func byProbeKey(entries []entry) map[string]float64 {
	out := make(map[string]float64, len(entries))
	for _, e := range entries {
		out[probeKey(e.labels)] = e.value
	}
	return out
}

func newProbe(success entry, series probeSeries) availability.Probe {
	key := probeKey(success.labels)
	kind := availability.ProbeKind(strings.TrimPrefix(success.labels["job"], "probe_"))
	probe := availability.Probe{
		ID:              probeID(kind, success.labels),
		Kind:            kind,
		Name:            success.labels["instance"],
		Target:          success.labels["target"],
		Up:              success.value == 1,
		Duration:        series.duration[key],
		Availability24h: probeRatio(series.day, key),
		Availability7d:  probeRatio(series.week, key),
		Availability30d: probeRatio(series.month, key),
	}
	if vmid, err := strconv.Atoi(success.labels["vmid"]); err == nil {
		probe.VMID = &vmid
	}
	if code, ok := series.status[key]; ok {
		n := int(code)
		probe.HTTPStatus = &n
	}
	return probe
}

func probeRatio(values map[string]float64, key string) *float64 {
	if v, ok := values[key]; ok {
		return &v
	}
	return nil
}
