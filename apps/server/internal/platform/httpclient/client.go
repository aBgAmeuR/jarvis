package httpclient

import (
	"net/http"
	"time"
)

type Options struct {
	CAFile string
	// CacheTTL of zero disables the cache.
	CacheTTL time.Duration
	// Timeout bounds an upstream call shared through the cache.
	Timeout      time.Duration
	MaxBodyBytes int64
	Wrap         []func(http.RoundTripper) http.RoundTripper
}

// New builds a client with no overall timeout: callers bound each call with a context.
func New(opts Options) (*http.Client, error) {
	base, err := newTransport(opts.CAFile)
	if err != nil {
		return nil, err
	}
	var rt http.RoundTripper = base
	for _, wrap := range opts.Wrap {
		rt = wrap(rt)
	}
	if opts.CacheTTL > 0 {
		rt = newCacheTransport(rt, opts.CacheTTL, opts.Timeout, opts.MaxBodyBytes)
	}
	return &http.Client{Transport: rt}, nil
}
