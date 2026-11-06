package services

type Series string

const (
	RequestsPerSecond Series = "requestsPerSecond"
	ErrorRatio        Series = "errorRatio"
	LatencyAvgSeconds Series = "latencyAvgSeconds"
	LatencyP95Seconds Series = "latencyP95Seconds"
	LatencyP99Seconds Series = "latencyP99Seconds"
)

func AllSeries() []Series {
	return []Series{
		RequestsPerSecond, ErrorRatio,
		LatencyAvgSeconds, LatencyP95Seconds, LatencyP99Seconds,
	}
}
