package availability

type ProbeSeries string

const (
	ProbeDuration    ProbeSeries = "durationSeconds"
	ProbeDurationAvg ProbeSeries = "durationAvgSeconds"
	ProbeDurationP99 ProbeSeries = "durationP99Seconds"
	ProbeSuccess     ProbeSeries = "success"
)

func AllProbeSeries() []ProbeSeries {
	return []ProbeSeries{
		ProbeDuration, ProbeDurationAvg, ProbeDurationP99, ProbeSuccess,
	}
}
