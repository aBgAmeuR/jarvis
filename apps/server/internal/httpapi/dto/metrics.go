package dto

import (
	"time"

	"jarvis/server/internal/domain"
)

type SeriesBody struct {
	Name       string       `json:"name" doc:"Series name." example:"cpuUsageRatio"`
	Unit       string       `json:"unit" doc:"Unit of the values." enum:"ratio,bytes,bytesPerSecond,load,celsius,seconds,perSecond" example:"ratio"`
	Query      string       `json:"query" doc:"PromQL that produced the series, for a Grafana Explore link." example:"max(pve_cpu_usage_ratio{id=\"node/pve\"})"`
	Timestamps []time.Time  `json:"timestamps" nullable:"false" doc:"Sample times, RFC 3339 in UTC. A gap is a missing timestamp."`
	Values     []float64    `json:"values" nullable:"false" doc:"Sample values, one per timestamp."`
	Summary    *SummaryBody `json:"summary,omitempty" doc:"Average and peak over the whole window. Absent on series without one, or with no sample."`
}

type SummaryBody struct {
	Avg  float64 `json:"avg" doc:"Average over the window." example:"0.12"`
	Peak float64 `json:"peak" doc:"Highest value over the window." example:"0.64"`
}

func newSummaryBody(s *domain.Summary) *SummaryBody {
	if s == nil {
		return nil
	}
	return &SummaryBody{Avg: s.Avg, Peak: s.Peak}
}

type MetricsBody struct {
	Start  time.Time    `json:"start" doc:"Start of the window actually read, aligned on its step." example:"2026-10-30T10:00:00Z"`
	End    time.Time    `json:"end" doc:"End of the window actually read, aligned on its step." example:"2026-10-31T10:00:00Z"`
	Series []SeriesBody `json:"series" nullable:"false" doc:"The requested series."`
}

func NewMetricsBody(ts domain.TimeSeries) MetricsBody {
	body := MetricsBody{Start: ts.Window.Start, End: ts.Window.End, Series: make([]SeriesBody, 0, len(ts.Series))}
	for _, s := range ts.Series {
		body.Series = append(body.Series, SeriesBody{Name: s.Name, Unit: string(s.Unit), Query: s.Query, Timestamps: s.Timestamps, Values: s.Values, Summary: newSummaryBody(s.Summary)})
	}
	return body
}
