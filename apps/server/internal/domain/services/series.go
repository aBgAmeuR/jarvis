package services

type Series string

const (
	RequestsPerSecond Series = "requestsPerSecond"
	ErrorRatio        Series = "errorRatio"
	LatencyP95Seconds Series = "latencyP95Seconds"
)

func AllSeries() []Series { return []Series{RequestsPerSecond, ErrorRatio, LatencyP95Seconds} }
