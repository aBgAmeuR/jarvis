package webui

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

// The SPA build is copied into static/ by the Dockerfile. Without it, only the API answers.
//
//go:embed all:static
var assets embed.FS

const (
	root  = "static/"
	shell = root + "_shell.html"

	// The shell carries inline scripts and styles from TanStack Start.
	csp = "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; " +
		"img-src 'self' data:; font-src 'self'; frame-ancestors 'none'"
)

// Handler serves the embedded SPA files and falls back to the shell for client routes.
func Handler() http.Handler {
	files := http.FileServerFS(assets)
	serve := func(w http.ResponseWriter, r *http.Request, file string) {
		req := r.Clone(r.Context())
		req.URL.Path = "/" + file
		files.ServeHTTP(w, req)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		h := w.Header()
		h.Set("Content-Security-Policy", csp)

		name := path.Clean(r.URL.Path)
		file := root + strings.TrimPrefix(name, "/")
		if info, err := fs.Stat(assets, file); err == nil && !info.IsDir() && !strings.HasPrefix(path.Base(name), ".") {
			if strings.HasPrefix(name, "/assets/") {
				h.Set("Cache-Control", "public, max-age=31536000, immutable")
			}
			serve(w, r, file)
			return
		}
		h.Set("Cache-Control", "no-cache")
		serve(w, r, shell)
	})
}
