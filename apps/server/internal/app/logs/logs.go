package logs

import "time"

const (
	DefaultLimit = 100
	MaxLimit     = 1000
)

type Logs struct {
	source LogSource
	now    func() time.Time
}

func NewLogs(source LogSource) *Logs {
	return &Logs{source: source, now: time.Now}
}
