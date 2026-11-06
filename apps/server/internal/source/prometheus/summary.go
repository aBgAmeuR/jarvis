package prometheus

import (
	"context"
	"errors"
	"strconv"
	"sync"

	"jarvis/server/internal/domain"
)

func (c *Client) querySeries(ctx context.Context, q query, w domain.Window) (domain.Series, error) {
	if q.samples == "" {
		return c.queryRange(ctx, q.expr, q.unit, q.name, w)
	}
	var (
		series               domain.Series
		summary              domain.Summary
		found                bool
		rangeErr, summaryErr error
		wg                   sync.WaitGroup
	)
	wg.Go(func() { series, rangeErr = c.queryRange(ctx, q.expr, q.unit, q.name, w) })
	wg.Go(func() { summary, found, summaryErr = c.querySummary(ctx, q, w) })
	wg.Wait()
	if rangeErr != nil {
		return domain.Series{}, rangeErr
	}
	// The summary only decorates the chart: when it fails, the series still shows without it.
	if found && summaryErr == nil {
		series.Summary = &summary
	}
	return series, nil
}

// querySummary takes the peak from every raw sample, so a spike between two steps still counts.
// The average reads the chart line at its step, which is cheap and close enough for a mean.
func (c *Client) querySummary(ctx context.Context, q query, w domain.Window) (domain.Summary, bool, error) {
	span := strconv.FormatInt(int64(w.End.Sub(w.Start).Seconds()), 10) + "s"
	step := strconv.FormatInt(int64(w.Step.Seconds()), 10) + "s"
	var (
		avg, peak       []entry
		avgErr, peakErr error
		wg              sync.WaitGroup
	)
	wg.Go(func() { avg, avgErr = c.queryVectorAt(ctx, "avg_over_time(("+q.expr+")["+span+":"+step+"])", w.End) })
	wg.Go(func() { peak, peakErr = c.queryVectorAt(ctx, "max(max_over_time("+q.samples+"["+span+"]))", w.End) })
	wg.Wait()
	if err := errors.Join(avgErr, peakErr); err != nil {
		return domain.Summary{}, false, err
	}
	if len(avg) == 0 || len(peak) == 0 {
		return domain.Summary{}, false, nil
	}
	return domain.Summary{Avg: avg[0].value, Peak: peak[0].value}, true, nil
}
