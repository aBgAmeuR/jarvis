package logs

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"

	applogs "jarvis/server/internal/app/logs"
	domainlogs "jarvis/server/internal/domain/logs"
	"jarvis/server/internal/httpapi/problem"
)

const path = "/api/v1/logs"

type searchLogsInput struct {
	Guest   int       `query:"guest" doc:"Only lines of this guest VMID." example:"116"`
	Service string    `query:"service" doc:"Only lines of this unit or container." example:"blocky"`
	Level   string    `query:"level" enum:"emerg,alert,crit,error,warning,notice,info,debug" doc:"Only lines of this journal priority."`
	Q       string    `query:"q" doc:"Only lines containing this text, case-sensitive." example:"timeout"`
	Start   time.Time `query:"start" doc:"Oldest instant to read, RFC 3339. Defaults to 24 hours before end." example:"2026-10-30T10:00:00Z"`
	End     time.Time `query:"end" doc:"Newest instant to read, exclusive, RFC 3339. Defaults to now." example:"2026-10-31T10:00:00Z"`
	Limit   int       `query:"limit" minimum:"1" maximum:"1000" default:"100" doc:"Maximum number of lines."`
	Cursor  string    `query:"cursor" doc:"Where the previous page stopped, from its Link header. Replaces end."`
}

type searchLogsOutput struct {
	Link string `header:"Link" doc:"The URL of the next page with rel=next, when more lines may follow."`
	Body struct {
		Items []entryBody `json:"items" nullable:"false" doc:"Log lines, newest first."`
	}
}

func registerSearchLogs(api huma.API, problems *problem.Mapper, logs *applogs.Logs) {
	huma.Register(api, huma.Operation{
		OperationID: "searchLogs",
		Method:      http.MethodGet,
		Path:        "/logs",
		Summary:     "Search logs",
		Tags:        []string{"supervision"},
	}, func(ctx context.Context, in *searchLogsInput) (*searchLogsOutput, error) {
		end := in.End
		if in.Cursor != "" {
			nanos, err := strconv.ParseInt(in.Cursor, 10, 64)
			if err != nil {
				return nil, huma.Error422UnprocessableEntity("invalid cursor")
			}
			end = time.Unix(0, nanos)
		}
		page, err := logs.Search(ctx, domainlogs.Query{
			VMID: in.Guest, Service: in.Service, Level: in.Level, Contains: in.Q,
			Start: in.Start, End: end, Limit: in.Limit,
		})
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &searchLogsOutput{}
		out.Body.Items = make([]entryBody, 0, len(page.Entries))
		for _, e := range page.Entries {
			out.Body.Items = append(out.Body.Items, newEntryBody(e))
		}
		if page.Next != nil {
			out.Link = nextLink(in, *page.Next)
		}
		return out, nil
	})
}

func nextLink(in *searchLogsInput, next time.Time) string {
	params := url.Values{"cursor": {strconv.FormatInt(next.UnixNano(), 10)}, "limit": {strconv.Itoa(in.Limit)}}
	if in.Guest != 0 {
		params.Set("guest", strconv.Itoa(in.Guest))
	}
	for key, value := range map[string]string{"service": in.Service, "level": in.Level, "q": in.Q} {
		if value != "" {
			params.Set(key, value)
		}
	}
	if !in.Start.IsZero() {
		params.Set("start", in.Start.UTC().Format(time.RFC3339))
	}
	return "<" + path + "?" + params.Encode() + `>; rel="next"`
}
