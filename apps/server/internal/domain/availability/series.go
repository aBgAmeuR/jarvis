package availability

type ProbeSeries string

const (
	ProbeDuration ProbeSeries = "durationSeconds"
	ProbeSuccess  ProbeSeries = "success"
)

func AllProbeSeries() []ProbeSeries { return []ProbeSeries{ProbeDuration, ProbeSuccess} }
