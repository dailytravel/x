package migrations

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dailytravel/x/configuration/graph/model"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Place struct {
	Database *mongo.Database
	Model    *model.Place
}

// Create mongo collection
func (m *Place) Migrate() error {
	col := m.Database.Collection(m.Model.Collection())
	indexes, err := col.Indexes().List(context.Background())
	if err != nil {
		return err
	}

	indexNames := make(map[string]bool)
	for indexes.Next(context.Background()) {
		var index bson.M
		if err := indexes.Decode(&index); err != nil {
			return err
		}

		indexNames[index["name"].(string)] = true
	}

	for _, index := range m.Model.Index() {
		keys := index.Keys
		if keys != nil {
			indexName := ""
			for _, key := range keys.(bson.D) {
				indexName = key.Key
				break
			}

			if !indexNames[indexName] {
				if _, err := col.Indexes().CreateOne(context.Background(), index); err != nil {
					return err
				}
			}
		}
	}

	// check if the "countries" collection is empty
	filter := bson.D{}
	if err := col.FindOne(context.Background(), filter).Err(); err != nil {
		if err == mongo.ErrNoDocuments {

			// Get the current working directory
			currentDir, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			// Check if the CSV file exists before opening it
			filePath := filepath.Join(currentDir, "configuration", "assets", "countries.csv")
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				log.Fatalf("CSV file does not exist: %s", filePath)
			}

			// Open the CSV file
			f, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			// Create a new CSV reader
			r := csv.NewReader(f)

			// Skip the header row
			if _, err := r.Read(); err != nil {
				log.Fatal(err)
			}

			// Read the rest of the rows
			// Loop over the CSV records
			for {
				// Read the next record
				row, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}

				// Create a new Place object
				defaultLocale := os.Getenv("DEFAULT_LOCALE")
				countryCode := strings.TrimSpace(strings.Split(row[2], "/")[0])
				country := &model.Place{
					Locale: defaultLocale,
					Type:   "country",
					Name:   map[string]interface{}{defaultLocale: row[0]},
					Slug:   slug.Make(row[0]),
					Status: "ACTIVE",
					Model: model.Model{
						Metadata: map[string]interface{}{
							"code":    countryCode,
							"dialing": row[1],
						},
					},
				}

				// Insert the country into the "countries" collection
				_, err = col.InsertOne(context.Background(), country)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	return nil
}
