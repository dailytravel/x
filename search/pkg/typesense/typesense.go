package typesense

import (
	"os"

	"github.com/typesense/typesense-go/typesense"
)

var (
	TS *typesense.Client
)

func ConnectTypesense() *typesense.Client {
	client := typesense.NewClient(
		typesense.WithServer(os.Getenv("TYPESENSE_HOST")),
		typesense.WithAPIKey(os.Getenv("TYPESENSE_API_KEY")),
	)

	return client
}
