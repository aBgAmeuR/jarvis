package health

import (
	"time"

	"jarvis/server/internal/domain"
)

type SourceStatus struct {
	Name      string
	Up        bool
	Kind      domain.SourceErrorKind
	CheckedAt time.Time
}
