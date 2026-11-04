package logs

import (
	"time"

	"jarvis/server/internal/domain"
)

type Count struct {
	Time  time.Time
	Level string
	N     int
}

type Volume struct {
	Window domain.Window
	Counts []Count
}
