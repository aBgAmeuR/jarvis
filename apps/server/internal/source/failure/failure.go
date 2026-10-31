package failure

import (
	"context"
	"errors"
	"net"
	"net/url"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/platform/httpclient"
)

func Wrap(source string, err error) error {
	if err == nil {
		return nil
	}
	return &domain.SourceError{Source: source, Kind: classify(err), Err: err}
}

func classify(err error) domain.SourceErrorKind {
	var (
		status *httpclient.StatusError
		decode *httpclient.DecodeError
		netErr net.Error
		urlErr *url.Error
		opErr  *net.OpError
		dnsErr *net.DNSError
	)
	switch {
	case errors.Is(err, context.DeadlineExceeded) || (errors.As(err, &netErr) && netErr.Timeout()):
		return domain.SourceTimeout
	case errors.As(err, &status):
		return domain.SourceBadStatus
	case errors.As(err, &decode):
		return domain.SourceInvalidResponse
	case errors.As(err, &urlErr), errors.As(err, &opErr), errors.As(err, &dnsErr):
		return domain.SourceUnreachable
	}
	return domain.SourceUnavailable
}
