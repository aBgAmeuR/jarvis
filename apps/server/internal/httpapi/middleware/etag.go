package middleware

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
)

const maxETagBody = 4 << 20

// A body over 4 MiB is streamed without an ETag.
func ETag(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (r.Method != http.MethodGet && r.Method != http.MethodHead) || !strings.HasPrefix(r.URL.Path, "/api/") {
			next.ServeHTTP(w, r)
			return
		}
		buffered := &etagWriter{ResponseWriter: w}
		next.ServeHTTP(buffered, r)
		buffered.finish(r)
	})
}

type etagWriter struct {
	http.ResponseWriter
	status      int
	body        bytes.Buffer
	passthrough bool
}

func (e *etagWriter) WriteHeader(code int) {
	if e.status == 0 {
		e.status = code
	}
	if e.passthrough {
		e.ResponseWriter.WriteHeader(code)
	}
}

func (e *etagWriter) Write(b []byte) (int, error) {
	if e.status == 0 {
		e.status = http.StatusOK
	}
	if e.passthrough {
		return e.ResponseWriter.Write(b)
	}
	if e.body.Len()+len(b) > maxETagBody {
		e.passthrough = true
		e.ResponseWriter.WriteHeader(e.status)
		if _, err := e.ResponseWriter.Write(e.body.Bytes()); err != nil {
			return 0, err
		}
		return e.ResponseWriter.Write(b)
	}
	return e.body.Write(b)
}

func (e *etagWriter) Unwrap() http.ResponseWriter { return e.ResponseWriter }

func (e *etagWriter) finish(r *http.Request) {
	if e.passthrough {
		return
	}
	if e.status == 0 {
		e.status = http.StatusOK
	}
	h := e.Header()
	if e.status == http.StatusOK {
		sum := sha256.Sum256(e.body.Bytes())
		etag := `"` + hex.EncodeToString(sum[:16]) + `"`
		h.Set("ETag", etag)
		h.Set("Cache-Control", "no-cache")
		if matchesETag(r.Header.Get("If-None-Match"), etag) {
			h.Del("Content-Length")
			h.Del("Content-Type")
			e.ResponseWriter.WriteHeader(http.StatusNotModified)
			return
		}
	}
	e.ResponseWriter.WriteHeader(e.status)
	_, _ = e.ResponseWriter.Write(e.body.Bytes())
}

func matchesETag(header, etag string) bool {
	for candidate := range strings.SplitSeq(header, ",") {
		candidate = strings.TrimPrefix(strings.TrimSpace(candidate), "W/")
		if candidate == "*" || candidate == etag {
			return true
		}
	}
	return false
}
