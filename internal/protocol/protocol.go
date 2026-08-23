package protocol

const (
	OpValidate  = "validate"
	OpProxmoxLs = "proxmox.ls"
)

type Request struct {
	Op string `json:"op"`
}

type Response struct {
	Error string `json:"error,omitempty"`
}

// Guest is one row of jarvis proxmox ls (CLI-ready).
type Guest struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Power string `json:"power"`
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
