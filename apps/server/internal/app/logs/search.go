package logs

import (
	"context"
	"fmt"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/logs"
)

type Page struct {
	Entries []logs.Entry
	Next    *time.Time
	// Query selects the same lines in Grafana Explore.
	Query string
}

// Search takes the Next of the previous page as end to read the following one.
func (l *Logs) Search(ctx context.Context, q logs.Query) (Page, error) {
	if q.Limit == 0 {
		q.Limit = DefaultLimit
	}
	if q.Limit < 0 || q.Limit > MaxLimit {
		return Page{}, fmt.Errorf("%w: limit must be between 1 and %d", domain.ErrInvalid, MaxLimit)
	}
	if q.End.IsZero() {
		q.End = l.now()
	}
	if q.Start.IsZero() {
		q.Start = q.End.Add(-domain.DefaultWindow)
	}
	if !q.End.After(q.Start) {
		return Page{}, fmt.Errorf("%w: end must be after start", domain.ErrInvalid)
	}
	if q.End.Sub(q.Start) > domain.MaxWindow {
		return Page{}, fmt.Errorf("%w: the range is limited to %d days", domain.ErrInvalid, domain.MaxWindow/(24*time.Hour))
	}

	entries, err := l.source.Logs(ctx, q)
	if err != nil {
		return Page{}, fmt.Errorf("search logs: %w", err)
	}
	page := Page{Entries: entries, Query: l.source.LogQuery(q)}
	if len(entries) == q.Limit {
		next := entries[len(entries)-1].Timestamp
		page.Next = &next
	}
	return page, nil
}
