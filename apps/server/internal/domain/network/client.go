package network

import (
	"strings"
	"time"
)

type Lease struct {
	MAC       string
	IP        string
	ExpiresAt time.Time
}

// Usage covers the last 24 hours.
type Usage struct {
	Queries     float64
	Blocked     float64
	LastQueryAt *time.Time
}

type Client struct {
	ID             string
	MAC            string
	IP             string
	Name           string
	LeaseExpiresAt time.Time
	Queries        int64
	Blocked        int64
	LastQueryAt    *time.Time
}

// ClientID is the MAC address in lowercase with hyphens, as the API names a client.
func ClientID(mac string) string {
	return strings.ReplaceAll(strings.ToLower(mac), ":", "-")
}
