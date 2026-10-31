package availability

type ProbeKind string

const (
	ProbeICMP ProbeKind = "icmp"
	ProbeDNS  ProbeKind = "dns"
	ProbeHTTP ProbeKind = "http"
)

type Probe struct {
	ID         string
	Kind       ProbeKind
	Name       string
	Target     string
	VMID       *int
	Up         bool
	Duration   float64
	HTTPStatus *int
	// The availability ratios are nil when Prometheus has no sample over the period.
	Availability24h *float64
	Availability7d  *float64
	Availability30d *float64
}
