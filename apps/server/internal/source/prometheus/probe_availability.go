package prometheus

import (
	"context"
	"strconv"
	"sync"
	"time"

	"jarvis/server/internal/domain/availability"
	"jarvis/server/internal/source/failure"
)

const (
	day         = 24 * time.Hour
	bucketCount = 90
)

func (c *Client) ProbeAvailability(ctx context.Context, probe availability.Probe, now time.Time) ([]availability.Bucket, error) {
	selector := "{job=" + quote("probe_"+string(probe.Kind)) + ",target=" + quote(probe.Target) + "}"
	midnight := now.UTC().Truncate(day)
	first := midnight.Add(-bucketCount * day)
	sinceMidnight := max(now.Sub(midnight), time.Minute)

	var (
		days     []matrixSeries
		today    []entry
		daysErr  error
		todayErr error
	)
	var wg sync.WaitGroup
	wg.Go(func() {
		days, daysErr = c.queryMatrixBetween(ctx, "avg_over_time(probe_success"+selector+"[1d])", first.Add(day), midnight, day)
	})
	wg.Go(func() {
		window := strconv.FormatInt(int64(sinceMidnight.Seconds()), 10)
		today, todayErr = c.queryVector(ctx, "avg_over_time(probe_success"+selector+"["+window+"s])")
	})
	wg.Wait()
	for _, err := range []error{daysErr, todayErr} {
		if err != nil {
			return nil, failure.Wrap(name, err)
		}
	}

	// Each sample is stamped at the end of the day it covers.
	byEnd := map[int64]float64{}
	if len(days) > 0 {
		for _, s := range days[0].samples {
			byEnd[s.at.Unix()] = s.value
		}
	}
	buckets := make([]availability.Bucket, 0, bucketCount+1)
	for i := range bucketCount {
		start := first.Add(time.Duration(i) * day)
		bucket := availability.Bucket{Start: start}
		if v, ok := byEnd[start.Add(day).Unix()]; ok {
			bucket.Availability = &v
		}
		buckets = append(buckets, bucket)
	}
	current := availability.Bucket{Start: midnight}
	if len(today) > 0 {
		current.Availability = &today[0].value
	}
	return append(buckets, current), nil
}
