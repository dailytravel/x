package migrations

import (
	"context"

	"github.com/dailytravel/x/account/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	Database *mongo.Database
	Model    *model.Api
}

func (m *Api) Migrate() error {
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

	//check if client exists
	filter := bson.M{"name": "Default"}
	if err := col.FindOne(context.Background(), filter).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			//create default client
			client := &model.Api{
				Name:       "Default",
				Identifier: "https://api.trip.express/graphql",
				Algorithm:  "SHA256",
				Expiration: 86400,
				Status:     "ACTIVE",
			}

			if _, err := col.InsertOne(context.Background(), client); err != nil {
				return err
			}
		}
	}

	return nil
}
