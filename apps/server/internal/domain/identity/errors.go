package identity

import "errors"

var (
	ErrUntrustedPeer = errors.New("request does not come from the trusted proxy")
	ErrNoUser        = errors.New("no user forwarded by the trusted proxy")
)
