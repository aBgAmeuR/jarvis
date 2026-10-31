package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/platform/config"
	"jarvis/server/internal/platform/httpserver"
	"jarvis/server/internal/platform/telemetry"
)

const (
	serviceName    = "jarvis-api"
	telemetryFlush = 5 * time.Second
)

func Run(ctx context.Context, cfg config.Config, meta domain.Meta, logger *slog.Logger) error {
	flushTraces, err := telemetry.SetupTracing(ctx, cfg.OTLPEndpoint, serviceName, meta.Version)
	if err != nil {
		return fmt.Errorf("set up tracing: %w", err)
	}
	defer func() {
		flushCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), telemetryFlush)
		defer cancel()
		if err := flushTraces(flushCtx); err != nil {
			logger.Warn("flush traces", "error", err)
		}
	}()

	metrics := telemetry.NewMetrics(meta.Version, meta.Commit)
	httpClients, err := newClients(cfg, metrics)
	if err != nil {
		return fmt.Errorf("build http clients: %w", err)
	}
	deps := newDeps(cfg, meta, logger, newSources(cfg, httpClients))

	logger.InfoContext(ctx, "listening", "addr", cfg.ListenAddr)
	return httpserver.Run(ctx, cfg.ListenAddr, newHandler(logger, metrics, deps))
}
