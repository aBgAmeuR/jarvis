package prometheus

import (
	"context"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
	"jarvis/server/internal/source/failure"
)

// An alert that started before the window starts at its first sample.
func (c *Client) ResolvedAlerts(ctx context.Context, w domain.Window) ([]availability.Alert, error) {
	matrix, err := c.queryMatrix(ctx, `ALERTS{alertstate="firing"}`, w)
	if err != nil {
		return nil, failure.Wrap(name, err)
	}
	var alerts []availability.Alert
	for _, series := range matrix {
		labels := make(map[string]string, len(series.labels))
		for k, v := range series.labels {
			if k != "__name__" && k != "alertstate" {
				labels[k] = v
			}
		}
		for _, episode := range episodes(series.samples, w.Step) {
			if !episode.end.Before(w.End.Add(-2 * w.Step)) {
				continue
			}
			ends := episode.end.Add(w.Step)
			alerts = append(alerts, availability.Alert{
				Fingerprint: fingerprint(labels),
				Name:        labels["alertname"],
				Severity:    domain.ParseSeverity(labels["severity"]),
				State:       availability.AlertResolved,
				StartsAt:    episode.start,
				EndsAt:      &ends,
				Labels:      labels,
				Annotations: map[string]string{},
			})
		}
	}
	return alerts, nil
}

type episode struct{ start, end time.Time }

func episodes(samples []sample, step time.Duration) []episode {
	var out []episode
	for _, s := range samples {
		if n := len(out); n > 0 && s.at.Sub(out[n-1].end) <= 2*step {
			out[n-1].end = s.at
			continue
		}
		out = append(out, episode{start: s.at, end: s.at})
	}
	return out
}
