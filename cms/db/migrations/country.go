package migrations

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dailytravel/x/cms/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Country struct {
	Database *mongo.Database
	Model    *model.Country
}

// Create mongo collection
func (m *Country) Migrate() error {
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

	// check if certificate exists
	filter := bson.D{}
	if err := col.FindOne(context.Background(), filter).Err(); err != nil {
		if err == mongo.ErrNoDocuments {

			// Get the current working directory
			currentDir, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			// Check if the CSV file exists before opening it
			filePath := filepath.Join(currentDir, "cms", "assets", "countries.csv")
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

				// Create a new Country object
				country := &model.Country{
					Locale:  os.Getenv("DEFAULT_LOCALE"),
					Name:    map[string]interface{}{os.Getenv("DEFAULT_LOCALE"): row[0]},
					Dialing: row[1],
					Code:    strings.TrimSpace(strings.Split(row[2], "/")[0]),
				}

				// Insert the country into the "countries" collection
				_, err = col.InsertOne(context.Background(), country)
				if err != nil {
					log.Fatal(err)
				}
			}

		} else {
			return err
		}
	}

	return nil
}
