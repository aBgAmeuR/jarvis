package availability

import "time"

// Availability is nil when the probe has no sample that day.
type Bucket struct {
	Start        time.Time
	Availability *float64
}
