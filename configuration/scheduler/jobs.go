package scheduler

import (
	"context"
	"log"

	"github.com/dailytravel/x/configuration/graph/model"
	"github.com/dailytravel/x/configuration/pkg/database"
	"github.com/robfig/cron"
	"github.com/typesense/typesense-go/typesense/api"
	"go.mongodb.org/mongo-driver/bson"
)

func SyncPlacesJob() {
	c := cron.New()
	// c.AddFunc("@daily", func() {
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Starting synchronization...")
		syncPlacesToTypesense()
		log.Println("Synchronization complete.")
	})
	c.Start()
}

func syncPlacesToTypesense() {
	ctx := context.Background()

	placeModel := &model.Place{}
	if _, err := database.Client.Collection(placeModel.Collection()).Retrieve(); err != nil {
		// Assuming a 'Schema' method on model.Place returns the required collection schema for Typesense
		schema, ok := placeModel.Schema().(*api.CollectionSchema)
		if !ok {
			log.Println("Error casting schema to *api.CollectionSchema")
			return
		}
		if _, err := database.Client.Collections().Create(schema); err != nil {
			log.Println("Error creating collection:", err)
			return
		}
	}

	// 1. Get a cursor for all places from MongoDB
	cursor, err := database.Database.Collection("places").Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error fetching places:", err)
		return
	}
	defer cursor.Close(ctx)

	// Assuming `places` collection in Typesense
	collection := database.Client.Collection("places")

	// Helper function for upserting document
	upsertDocument := func(document map[string]interface{}) {
		if _, err := collection.Documents().Upsert(document); err != nil {
			log.Printf("Error upserting document %s: %v\n", document["id"], err)
		}
	}

	// Iterate through each place
	for cursor.Next(ctx) {
		var place model.Place
		if err := cursor.Decode(&place); err != nil {
			log.Println("Error decoding place:", err)
			continue
		}

		document := place.Document() // Assuming you have a Document method as before

		// 2. Check if place exists in Typesense
		tsDocument, err := collection.Document(place.ID.Hex()).Retrieve()

		// If there's an error (like document doesn't exist) or the update timestamps differ, upsert
		if err != nil || tsDocument["updated"].(string) != document["updated"].(string) {
			upsertDocument(document)
		}
	}

	// Check if cursor encountered any error during iteration
	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
	}
}
