package migrations

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dailytravel/x/configuration/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Currency struct {
	Database *mongo.Database
	Model    *model.Currency
}

// Create mongo collection
func (m *Currency) Migrate() error {
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
			filePath := filepath.Join(currentDir, "configuration", "assets", "currencies.csv")
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				log.Fatalf("CSV file does not exist: %s", filePath)
			}

			// Open the CSV file
			f, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			// Read CSV file
			reader := csv.NewReader(f)
			reader.FieldsPerRecord = -1 // Allow variable number of fields per record
			rows, err := reader.ReadAll()
			if err != nil {
				log.Fatal(err)
			}

			// Convert CSV rows to Currency struct and insert into MongoDB
			for i, row := range rows {
				if i == 0 {
					continue // Skip header row
				}

				// Create a new Currency object
				currency := &model.Currency{
					Code:     row[0],
					Locale:   os.Getenv("DEFAULT_LOCALE"),
					Name:     map[string]interface{}{os.Getenv("DEFAULT_LOCALE"): row[1]},
					Symbol:   row[2],
					Decimal:  row[5],
					Thousand: row[6],
				}

				rate, err := strconv.ParseFloat(row[3], 64)
				if err == nil {
					currency.Rate = rate
				}

				precision, err := strconv.Atoi(row[4])
				if err == nil {
					currency.Precision = precision
				}

				order, err := strconv.Atoi(row[7])
				if err == nil {
					currency.Order = order
				}

				// Insert the currency into the "currencies" collection
				_, err = col.InsertOne(context.Background(), currency)
				if err != nil {
					log.Fatal(err)
				}
			}

		}
	}

	return nil
}
