package middleware

import (
	"context"
	"net/http"
	"time"
)

type Observation struct {
	Method   string
	Path     string
	Route    string
	Status   int
	Bytes    int
	Duration time.Duration
}

type Observer func(ctx context.Context, o Observation)

// Place it outside middleware that replaces the request, so the matched route stays visible.
func Observe(observers ...Observer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rec := &recorder{ResponseWriter: w}
			next.ServeHTTP(rec, r)
			if rec.status == 0 {
				rec.status = http.StatusOK
			}
			o := Observation{
				Method:   r.Method,
				Path:     r.URL.Path,
				Route:    r.Pattern,
				Status:   rec.status,
				Bytes:    rec.bytes,
				Duration: time.Since(start),
			}
			for _, observe := range observers {
				observe(r.Context(), o)
			}
		})
	}
}
