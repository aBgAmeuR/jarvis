package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"

	"jarvis/server/internal/config"
	"jarvis/server/internal/domain"
	"jarvis/server/internal/httpapi"
	"jarvis/server/internal/httpserver"
)

var (
	version   = "dev"
	commit    = "none"
	buildDate = "unknown"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	if err := run(logger); err != nil {
		logger.Error("exit", "error", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	logger.Info("config", "config", cfg)

	mux := http.NewServeMux()
	httpapi.NewAPI(mux, domain.Meta{Version: version, Commit: commit, BuildDate: buildDate})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	logger.Info("listening", "addr", cfg.ListenAddr)
	return httpserver.Run(ctx, cfg.ListenAddr, mux)
}
