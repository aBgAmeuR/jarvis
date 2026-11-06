package network

// DNS covers the last 24 hours.
type DNS struct {
	Queries                int64
	Blocked                int64
	Cached                 int64
	UpstreamLatencySeconds float64
	QueriesPerSecond       float64
}

type DNSSeries string

const (
	DNSQueriesPerSecond   DNSSeries = "queriesPerSecond"
	DNSBlockedPerSecond   DNSSeries = "blockedPerSecond"
	DNSUpstreamLatency    DNSSeries = "upstreamLatencySeconds"
	DNSUpstreamLatencyP99 DNSSeries = "upstreamLatencyP99Seconds"
)

func AllDNSSeries() []DNSSeries {
	return []DNSSeries{
		DNSQueriesPerSecond, DNSBlockedPerSecond,
		DNSUpstreamLatency, DNSUpstreamLatencyP99,
	}
}

type ClientSeries string

const (
	ClientQueriesPerSecond ClientSeries = "queriesPerSecond"
	ClientBlockedPerSecond ClientSeries = "blockedPerSecond"
)

func AllClientSeries() []ClientSeries {
	return []ClientSeries{ClientQueriesPerSecond, ClientBlockedPerSecond}
}
