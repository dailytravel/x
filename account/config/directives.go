package config

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dailytravel/x/account/auth"
	"github.com/dailytravel/x/account/graph"
)

var (
	ErrAPIKeyRequired   = errors.New("API key is required")
	ErrNotAuthenticated = errors.New("not authenticated")
	ErrAccessDenied     = errors.New("access denied")
	ErrMissingRole      = errors.New("access denied: missing role")
	ErrMissingScope     = errors.New("access denied: missing scope")
)

func Directives(c *graph.Config) {
	c.Directives.Api = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		return next(ctx)
	}

	c.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver, requires []*string) (interface{}, error) {
		claims := auth.Auth(ctx)
		if claims == "" {
			return nil, ErrNotAuthenticated
		}

		if len(requires) == 0 {
			return next(ctx)
		}

		// for _, require := range requires {
		// 	for _, role := range strings.Split(claims["roles"], " ") {
		// 		if role == *require {
		// 			return next(ctx)
		// 		}
		// 	}
		// }

		return nil, ErrMissingRole
	}

	c.Directives.HasScope = func(ctx context.Context, obj interface{}, next graphql.Resolver, scope []string) (interface{}, error) {
		claims := auth.Auth(ctx)
		if claims == "" {
			return nil, ErrNotAuthenticated
		}

		if len(scope) == 0 {
			return next(ctx)
		}

		// for _, s := range scope {
		// 	for _, r := range strings.Split(claims["scopes"], " ") {
		// 		if r == s {
		// 			return next(ctx)
		// 		}
		// 	}
		// }

		return nil, ErrMissingScope
	}
}
