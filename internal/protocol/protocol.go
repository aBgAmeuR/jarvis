package protocol

import "github.com/abgameur/jarvis/internal/config"

type Request struct {
	Op string `json:"op"`
}

type Response struct {
	Services []config.Service `json:"services,omitempty"`
	Error    string           `json:"error,omitempty"`
}
