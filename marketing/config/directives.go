package config

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dailytravel/x/marketing/graph"
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
		return nil, ErrMissingRole
	}

	c.Directives.HasScope = func(ctx context.Context, obj interface{}, next graphql.Resolver, scope []string) (interface{}, error) {
		return nil, ErrMissingScope
	}
}
