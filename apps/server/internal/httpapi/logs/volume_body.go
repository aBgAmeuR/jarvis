package logs

import (
	"time"

	domainlogs "jarvis/server/internal/domain/logs"
)

type volumeCountBody struct {
	Time  time.Time `json:"time" doc:"Start of the bucket, RFC 3339 in UTC." example:"2026-10-31T10:00:00Z"`
	Level string    `json:"level,omitempty" doc:"Journal priority. Absent when the stream has none." example:"info"`
	Count int       `json:"count" doc:"Lines in the bucket." example:"42"`
}

func newVolumeCountBody(c domainlogs.Count) volumeCountBody {
	return volumeCountBody{Time: c.Time, Level: c.Level, Count: c.N}
}
