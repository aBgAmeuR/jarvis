package identity

import (
	"net/netip"
	"strings"

	"jarvis/server/internal/domain/identity"
)

// Resolver trusts the Remote-* headers from one proxy address only.
type Resolver struct {
	trustedProxy netip.Addr
}

func NewResolver(trustedProxy netip.Addr) *Resolver {
	return &Resolver{trustedProxy: trustedProxy.Unmap()}
}

func (r *Resolver) Resolve(peer netip.Addr, name, groups string) (identity.User, error) {
	// IPv4 peers can arrive as ::ffff:a.b.c.d.
	if peer.Unmap() != r.trustedProxy {
		return identity.User{}, identity.ErrUntrustedPeer
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return identity.User{}, identity.ErrNoUser
	}
	list := []string{}
	for g := range strings.SplitSeq(groups, ",") {
		if g = strings.TrimSpace(g); g != "" {
			list = append(list, g)
		}
	}
	return identity.User{Name: name, Groups: list}, nil
}
