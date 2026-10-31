package overview

type Status string

const (
	StatusOK       Status = "ok"
	StatusInfo     Status = "info"
	StatusWarning  Status = "warning"
	StatusCritical Status = "critical"
	StatusUnknown  Status = "unknown"
)
