package config

import (
	"context"
	"errors"
	"strings"

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
		apiKey := auth.APIKey(ctx)
		if apiKey == nil {
			return nil, ErrAPIKeyRequired
		}

		// Return next(ctx) without the need for an additional "else" block
		return next(ctx)
	}

	c.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []string) (interface{}, error) {
		claims := auth.Auth(ctx)
		if claims == nil {
			return nil, ErrNotAuthenticated
		}

		rolesClaim, ok := claims["roles"].(string)
		if !ok {
			return nil, ErrMissingRole
		}

		if len(roles) == 0 {
			return next(ctx)
		}

		requiredRoles := make(map[string]bool)
		for _, role := range roles {
			requiredRoles[strings.ToLower(role)] = true
		}

		roleArray := strings.Split(rolesClaim, " ")
		for _, role := range roleArray {
			if requiredRoles[strings.ToLower(role)] {
				return next(ctx)
			}
		}

		return nil, ErrMissingRole
	}

	c.Directives.HasScope = func(ctx context.Context, obj interface{}, next graphql.Resolver, scope []string) (interface{}, error) {
		claims := auth.Auth(ctx)
		if claims == nil {
			return nil, ErrNotAuthenticated
		}

		scopeClaim, ok := claims["scope"].(string)
		if !ok {
			return nil, ErrMissingScope
		}

		if len(scope) == 0 || scopeClaim == "*" {
			return next(ctx)
		}

		requiredScopes := make(map[string]bool)
		for _, s := range scope {
			requiredScopes[s] = true
		}

		scopeArray := strings.Split(scopeClaim, " ")
		for _, s := range scopeArray {
			if requiredScopes[s] {
				return next(ctx)
			}
		}

		return nil, ErrMissingScope
	}
}
