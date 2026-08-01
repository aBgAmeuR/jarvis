package protocol

const OpListServices = "list_services"

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
