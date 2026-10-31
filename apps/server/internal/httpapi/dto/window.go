package dto

import "time"

type WindowInput struct {
	Start time.Time `query:"start" doc:"Start of the window, RFC 3339. Defaults to 24 hours before end." example:"2026-10-30T10:00:00Z"`
	End   time.Time `query:"end" doc:"End of the window, RFC 3339. Defaults to now." example:"2026-10-31T10:00:00Z"`
}
