package migrations

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/dailytravel/x/configuration/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Option struct {
	Database *mongo.Database
	Model    *model.Option
}

// Create mongo collection
func (m *Option) Migrate() error {
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

	// check if the "options" collection is empty
	filter := bson.D{}
	if err := col.FindOne(context.Background(), filter).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			var data map[string]interface{}

			// Get the current working directory
			currentDir, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			// Check if the json file exists
			filePath := filepath.Join(currentDir, "configuration", "assets", "options.json")
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				log.Fatalf("JSON file does not exist: %s", filePath)
			}
			// Open the CSV file
			file, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			byteData, _ := ioutil.ReadAll(file)
			_ = json.Unmarshal(byteData, &data)

			var docs []interface{}

			for k, v := range data {
				doc := &model.Option{
					Name: k,
					Data: fmt.Sprintf("%v", v),
				}
				docs = append(docs, doc)
			}

			if _, err := col.InsertMany(context.Background(), docs); err != nil {
				panic(err)
			}
		}
	}

	return nil
}
