package database

import (
	"github.com/typesense/typesense-go/typesense"
)

var (
	Client *typesense.Client
)

func ConnectTypesense() *typesense.Client {
	client := typesense.NewClient(
		typesense.WithServer("http://localhost:8108"),
		typesense.WithAPIKey("5213473d89e548d14568ea01a2309f72"),
	)

	return client
}
