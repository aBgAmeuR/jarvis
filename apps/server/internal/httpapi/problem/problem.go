package problem

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"jarvis/server/internal/domain"
)

// Mapper turns the errors of the use cases into Huma errors.
type Mapper struct {
	logger *slog.Logger
}

func NewMapper(logger *slog.Logger) *Mapper {
	return &Mapper{logger: logger}
}

// Error maps a missing object to a 404 named after its kind, a source failure to a 502 naming the source,
// and a refused request to a 422. The detail of a failure goes to the logs.
func (m *Mapper) Error(ctx context.Context, err error) error {
	var notFound *domain.NotFoundError
	if errors.As(err, &notFound) {
		return huma.Error404NotFound(upperFirst(notFound.Kind) + " not found")
	}
	if errors.Is(err, domain.ErrNotFound) {
		return huma.Error404NotFound("Not found")
	}
	if errors.Is(err, domain.ErrInvalid) {
		return huma.Error422UnprocessableEntity(err.Error())
	}
	var sourceErr *domain.SourceError
	if errors.As(err, &sourceErr) {
		m.logger.ErrorContext(ctx, "data source failed", "source", sourceErr.Source, "kind", sourceErr.Kind, "error", err)
		return huma.Error502BadGateway(fmt.Sprintf("%s is unavailable", sourceErr.Source))
	}
	m.logger.ErrorContext(ctx, "request failed", "error", err)
	return huma.Error500InternalServerError("Internal server error")
}

func upperFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
