package logs

import (
	"time"

	domainlogs "jarvis/server/internal/domain/logs"
)

type entryBody struct {
	Timestamp time.Time `json:"timestamp" doc:"When the line was written, RFC 3339 in UTC." example:"2026-10-31T18:23:58.123456789Z"`
	Line      string    `json:"line" doc:"The log line." example:"query resolved"`
	Level     string    `json:"level,omitempty" doc:"Journal priority. Absent when the stream has none." example:"info"`
	Service   string    `json:"service" doc:"Unit or container the line comes from." example:"blocky"`
	Instance  string    `json:"instance,omitempty" doc:"Guest name." example:"dns"`
	VMID      *int      `json:"vmid,omitempty" doc:"Guest VMID. Absent when the stream has none." example:"116"`
}

func newEntryBody(e domainlogs.Entry) entryBody {
	body := entryBody{Timestamp: e.Timestamp, Line: e.Line, Level: e.Level, Service: e.Service, Instance: e.Instance}
	if e.VMID != 0 {
		body.VMID = &e.VMID
	}
	return body
}
