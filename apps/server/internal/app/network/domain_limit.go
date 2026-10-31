package network

import (
	"fmt"

	"jarvis/server/internal/domain"
)

const (
	DefaultDomains = 20
	MaxDomains     = 100
)

func domainLimit(limit int) (int, error) {
	if limit == 0 {
		return DefaultDomains, nil
	}
	if limit < 0 || limit > MaxDomains {
		return 0, fmt.Errorf("%w: limit must be between 1 and %d", domain.ErrInvalid, MaxDomains)
	}
	return limit, nil
}
