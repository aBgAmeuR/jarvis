// Package protocol defines the wire types shared by the CLI and jarvisd.
package protocol

const (
	OpValidate     = "validate"
	OpProxmoxLs    = "proxmox.ls"
	OpProxmoxStart = "proxmox.start"
	OpProxmoxStop  = "proxmox.stop"
)

type Request struct {
	Op   string `json:"op"`
	VMID int    `json:"vmid,omitempty"`
}

// Response carries the error of a failed unary op. Success replies embed
// their payload alongside an empty error.
type Response struct {
	Error string `json:"error,omitempty"`
}

const (
	GuestTypeLXC = "lxc"
	GuestTypeVM  = "vm"
)

// Guest is one row of jarvis proxmox ls (CLI-ready).
type Guest struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Power string `json:"power"`
}

// ProxmoxLsResult is the success payload of OpProxmoxLs.
type ProxmoxLsResult struct {
	Guests []Guest `json:"guests"`
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
