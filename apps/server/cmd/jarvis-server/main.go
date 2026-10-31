package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"jarvis/server/internal/bootstrap"
	"jarvis/server/internal/domain"
	"jarvis/server/internal/platform/config"
	"jarvis/server/internal/platform/logging"
)

var (
	version   = "dev"
	commit    = "none"
	buildDate = "unknown"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	logger := logging.New(os.Stderr, cfg.LogLevel)
	if err != nil {
		logger.Error("exit", "error", err)
		os.Exit(1)
	}
	logger.Info("config", "config", cfg)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	meta := domain.Meta{Version: version, Commit: commit, BuildDate: buildDate}
	if err := bootstrap.Run(ctx, cfg, meta, logger); err != nil {
		logger.Error("exit", "error", err)
		os.Exit(1)
	}
}
