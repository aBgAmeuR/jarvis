package httpclient

import (
	"context"
	"net/http"
	"sync"
	"time"
)

const (
	maxCacheEntries = 128
	maxCachedBytes  = 1 << 20
)

// cacheTransport shares one upstream call between identical GETs; it outlives its first caller.
type cacheTransport struct {
	next    http.RoundTripper
	ttl     time.Duration
	timeout time.Duration
	maxBody int64

	mu      sync.Mutex
	entries map[string]*cacheEntry
}

type cacheEntry struct {
	done    chan struct{}
	snap    *snapshot
	err     error
	expires time.Time
}

func newCacheTransport(next http.RoundTripper, ttl, timeout time.Duration, maxBody int64) *cacheTransport {
	return &cacheTransport{next: next, ttl: ttl, timeout: timeout, maxBody: maxBody, entries: map[string]*cacheEntry{}}
}

func (e *cacheEntry) usable(now time.Time) bool {
	select {
	case <-e.done:
		return e.err == nil && now.Before(e.expires)
	default:
		return true
	}
}

func (t *cacheTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method != http.MethodGet {
		return t.next.RoundTrip(req)
	}
	entry := t.entry(req)
	if entry == nil {
		return t.next.RoundTrip(req)
	}
	select {
	case <-entry.done:
	case <-req.Context().Done():
		return nil, context.Cause(req.Context())
	}
	if entry.err != nil {
		return nil, entry.err
	}
	return entry.snap.response(req), nil
}

// entry returns nil when the cache is full, so the caller goes straight upstream.
func (t *cacheTransport) entry(req *http.Request) *cacheEntry {
	key := req.URL.String()
	now := time.Now()

	t.mu.Lock()
	defer t.mu.Unlock()
	if e, ok := t.entries[key]; ok && e.usable(now) {
		return e
	}
	if len(t.entries) >= maxCacheEntries {
		t.sweep(now)
		if len(t.entries) >= maxCacheEntries {
			return nil
		}
	}
	e := &cacheEntry{done: make(chan struct{})}
	t.entries[key] = e
	go t.fill(key, e, req.Clone(context.WithoutCancel(req.Context())))
	return e
}

func (t *cacheTransport) sweep(now time.Time) {
	for key, e := range t.entries {
		if !e.usable(now) {
			delete(t.entries, key)
		}
	}
}

func (t *cacheTransport) fill(key string, e *cacheEntry, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), t.timeout)
	defer cancel()
	snap, err := t.fetch(req.WithContext(ctx))

	t.mu.Lock()
	e.snap, e.err = snap, err
	if err == nil && snap.storable(maxCachedBytes) {
		e.expires = time.Now().Add(t.ttl)
	} else {
		delete(t.entries, key)
	}
	t.mu.Unlock()
	close(e.done)
}

func (t *cacheTransport) fetch(req *http.Request) (*snapshot, error) {
	resp, err := t.next.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	return readSnapshot(resp, t.maxBody)
}
