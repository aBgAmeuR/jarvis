package healthcheck

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

const timeout = 3 * time.Second

// Run calls /healthz on the local server. The image has no shell or curl, so the binary checks itself.
func Run(ctx context.Context, listenAddr string) error {
	host, port, err := net.SplitHostPort(listenAddr)
	if err != nil {
		return fmt.Errorf("parse listen address: %w", err)
	}
	if host == "" {
		host = "127.0.0.1"
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://"+net.JoinHostPort(host, port)+"/healthz", http.NoBody)
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("call healthz: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("healthz answered %s", resp.Status)
	}
	return nil
}
