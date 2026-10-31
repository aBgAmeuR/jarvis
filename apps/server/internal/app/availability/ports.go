package availability

import (
	"context"
	"time"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/availability"
)

type AlertSource interface {
	Alerts(ctx context.Context) ([]availability.Alert, error)
}

type AlertHistorySource interface {
	ResolvedAlerts(ctx context.Context, w domain.Window) ([]availability.Alert, error)
}

type ProbeSource interface {
	Probes(ctx context.Context) ([]availability.Probe, error)
	ProbeSeries(ctx context.Context, probe availability.Probe, names []availability.ProbeSeries, w domain.Window) ([]domain.Series, error)
	ProbeAvailability(ctx context.Context, probe availability.Probe, now time.Time) ([]availability.Bucket, error)
}

type CertificateSource interface {
	Certificates(ctx context.Context) ([]availability.Certificate, error)
}

type SilenceSource interface {
	Silences(ctx context.Context) ([]availability.Silence, error)
}
