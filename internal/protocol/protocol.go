package protocol

const (
	OpListServices = "list_services"
	OpValidate     = "validate"
)

type Request struct {
	Op string `json:"op"`
}

type ServiceInfo struct {
	Name string `json:"name"`
}

type Response struct {
	Services []ServiceInfo `json:"services,omitempty"`
	Error    string        `json:"error,omitempty"`
}

// ValidateEvent is one NDJSON line from OpValidate.
type ValidateEvent struct {
	Layer string `json:"layer"`
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

const (
	LayerParse       = "parse"
	LayerRequired    = "required"
	LayerValues      = "values"
	LayerSecrets     = "secrets"
	LayerFingerprint = "fingerprint"
	LayerSSH         = "ssh"
	LayerGuests      = "guests"
)
