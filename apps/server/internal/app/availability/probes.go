package availability

import "time"

type Probes struct {
	source ProbeSource
	now    func() time.Time
}

func NewProbes(source ProbeSource) *Probes {
	return &Probes{source: source, now: time.Now}
}
