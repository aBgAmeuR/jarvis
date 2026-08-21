package protocol

const (
	OpValidate = "validate"
)

type Request struct {
	Op string `json:"op"`
}

type Response struct {
	Error string `json:"error,omitempty"`
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
