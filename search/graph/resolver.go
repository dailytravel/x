package graph

import (
	"github.com/typesense/typesense-go/typesense"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ts *typesense.Client
}

func NewResolver(client *typesense.Client) *Resolver {
	return &Resolver{
		ts: client,
	}
}
