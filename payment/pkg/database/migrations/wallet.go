package migrations

import (
	"context"

	"github.com/dailytravel/x/payment/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Wallet struct {
	Database *mongo.Database
	Model    *model.Wallet
}

// Create mongo collection
func (m *Wallet) Migrate() error {
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

	return nil
}
