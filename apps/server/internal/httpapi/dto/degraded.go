package dto

import "jarvis/server/internal/domain"

type DegradedBody struct {
	Source string `json:"source" doc:"Data source that failed." example:"alertmanager"`
	Reason string `json:"reason" doc:"How it failed." example:"Did not answer in time"`
}

func Reason(kind domain.SourceErrorKind) string {
	switch kind {
	case domain.SourceTimeout:
		return "Did not answer in time"
	case domain.SourceUnreachable:
		return "Unreachable"
	case domain.SourceBadStatus:
		return "Answered with an error status"
	case domain.SourceInvalidResponse:
		return "Answered with an unreadable response"
	default:
		return "Unavailable"
	}
}

func NewDegradedBodies(list []domain.Degraded) []DegradedBody {
	out := make([]DegradedBody, 0, len(list))
	for _, d := range list {
		out = append(out, DegradedBody{Source: d.Source, Reason: Reason(d.Kind)})
	}
	return out
}
