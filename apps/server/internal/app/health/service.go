package health

import (
	"context"
	"sync"
	"time"

	"jarvis/server/internal/domain"
	domainhealth "jarvis/server/internal/domain/health"
)

type Service struct {
	checkers []Checker
	timeout  time.Duration
}

func NewService(timeout time.Duration, checkers ...Checker) *Service {
	return &Service{checkers: checkers, timeout: timeout}
}

func (s *Service) Status(ctx context.Context) []domainhealth.SourceStatus {
	out := make([]domainhealth.SourceStatus, len(s.checkers))
	var wg sync.WaitGroup
	for i, checker := range s.checkers {
		wg.Go(func() { out[i] = s.check(ctx, checker) })
	}
	wg.Wait()
	return out
}

func (s *Service) check(ctx context.Context, checker Checker) domainhealth.SourceStatus {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	status := domainhealth.SourceStatus{Name: checker.Name(), Up: true, CheckedAt: time.Now().UTC()}
	if err := checker.Ping(ctx); err != nil {
		status.Up, status.Kind = false, domain.NewDegraded(checker.Name(), err).Kind
	}
	return status
}

func (s *Service) Ready(ctx context.Context) bool {
	for _, status := range s.Status(ctx) {
		if !status.Up {
			return false
		}
	}
	return true
}
