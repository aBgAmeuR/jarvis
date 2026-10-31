package domain

import (
	"fmt"
	"time"
)

const (
	DefaultWindow = 24 * time.Hour
	MaxWindow     = 90 * 24 * time.Hour

	targetPoints = 720
	minRate      = time.Minute
)

// Finest first.
var steps = [...]time.Duration{
	15 * time.Second, 30 * time.Second, time.Minute, 2 * time.Minute,
	5 * time.Minute, 15 * time.Minute, 30 * time.Minute, time.Hour,
}

type Window struct {
	Start time.Time
	End   time.Time
	Step  time.Duration
}

// NewWindow aligns both ends on a step that keeps the range under 720 points, so identical requests share data.
func NewWindow(start, end, now time.Time) (Window, error) {
	if end.IsZero() {
		end = now
	}
	if start.IsZero() {
		start = end.Add(-DefaultWindow)
	}
	span := end.Sub(start)
	switch {
	case span <= 0:
		return Window{}, fmt.Errorf("%w: end must be after start", ErrInvalid)
	case span > MaxWindow:
		return Window{}, fmt.Errorf("%w: the range is limited to %d days", ErrInvalid, MaxWindow/(24*time.Hour))
	}

	step := stepFor(span)
	w := Window{Start: start.UTC().Truncate(step), End: end.UTC().Truncate(step), Step: step}
	if !w.End.After(w.Start) {
		return Window{}, fmt.Errorf("%w: the range is shorter than one step of %s", ErrInvalid, step)
	}
	return w, nil
}

func stepFor(span time.Duration) time.Duration {
	want := span / targetPoints
	for _, step := range steps {
		if step >= want {
			return step
		}
	}
	return steps[len(steps)-1]
}

// RateInterval is the step, and four scrapes at least.
func (w Window) RateInterval() time.Duration {
	return max(w.Step, minRate)
}
