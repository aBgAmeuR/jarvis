package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/netip"

	"github.com/danielgtaylor/huma/v2"

	appidentity "jarvis/server/internal/app/identity"
	domainidentity "jarvis/server/internal/domain/identity"
)

type userKey struct{}

func UserFrom(ctx context.Context) (domainidentity.User, bool) {
	user, ok := ctx.Value(userKey{}).(domainidentity.User)
	return user, ok
}

// The cause is logged, never sent, and never includes a header value.
func Identity(api huma.API, logger *slog.Logger, resolver *appidentity.Resolver) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		user, err := resolve(ctx, resolver)
		if err != nil {
			logger.WarnContext(ctx.Context(), "identity refused", "peer", ctx.RemoteAddr(), "reason", err.Error())
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}
		next(huma.WithValue(ctx, userKey{}, user))
	}
}

func resolve(ctx huma.Context, resolver *appidentity.Resolver) (domainidentity.User, error) {
	peer, err := netip.ParseAddrPort(ctx.RemoteAddr())
	if err != nil {
		return domainidentity.User{}, fmt.Errorf("parse peer address: %w", err)
	}
	return resolver.Resolve(peer.Addr(), ctx.Header("Remote-User"), ctx.Header("Remote-Groups"))
}
