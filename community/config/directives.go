package config

import (
	"context"
	"errors"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dailytravel/x/community/graph"
	"github.com/dailytravel/x/community/pkg/auth"
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

		userRolesInterface, ok := claims["roles"]
		if !ok {
			return nil, ErrMissingRole
		}

		// Check for the presence of a role without converting to a map
		if len(roles) == 0 || hasRole(userRolesInterface, roles) {
			return next(ctx)
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

// Helper function to check if userRoles contains any of the provided roles
func hasRole(userRolesInterface interface{}, roles []string) bool {
	userRoles, ok := userRolesInterface.([]interface{})
	if !ok {
		return false
	}

	for _, roleInterface := range userRoles {
		role, ok := roleInterface.(string)
		if !ok {
			continue
		}

		for _, r := range roles {
			if strings.EqualFold(role, r) {
				return true
			}
		}
	}
	return false
}
