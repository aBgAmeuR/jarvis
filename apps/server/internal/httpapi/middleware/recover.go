package middleware

import (
	"errors"
	"log/slog"
	"net/http"
	"runtime/debug"
)

const internalErrorBody = `{"title":"Internal Server Error","status":500}`

func Recover(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer recoverRequest(logger, w, r)
			next.ServeHTTP(w, r)
		})
	}
}

func recoverRequest(logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	v := recover()
	if v == nil {
		return
	}
	if err, ok := v.(error); ok && errors.Is(err, http.ErrAbortHandler) {
		panic(v)
	}
	logger.ErrorContext(r.Context(), "panic serving request", "panic", v, "stack", string(debug.Stack()))
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(internalErrorBody))
}
