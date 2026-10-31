package availability

import "time"

type Alerts struct {
	source  AlertSource
	history AlertHistorySource
	now     func() time.Time
}

func NewAlerts(source AlertSource, history AlertHistorySource) *Alerts {
	return &Alerts{source: source, history: history, now: time.Now}
}
