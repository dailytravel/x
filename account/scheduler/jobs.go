package scheduler

import (
	"context"
	"log"

	"github.com/dailytravel/x/account/graph/model"
	"github.com/dailytravel/x/account/pkg/database"
	"github.com/robfig/cron"
	"github.com/typesense/typesense-go/typesense/api"
	"go.mongodb.org/mongo-driver/bson"
)

func SyncUsersJob() {
	c := cron.New()
	// c.AddFunc("@daily", func() {
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Starting synchronization...")
		syncUsersToTypesense()
		log.Println("Synchronization complete.")
	})
	c.Start()
}

func syncUsersToTypesense() {
	ctx := context.Background()

	item := &model.User{}
	if _, err := database.Client.Collection(item.Collection()).Retrieve(); err != nil {
		// Assuming a 'Schema' method on model.User returns the required collection schema for Typesense
		schema, ok := item.Schema().(*api.CollectionSchema)
		if !ok {
			log.Println("Error casting schema to *api.CollectionSchema")
			return
		}
		if _, err := database.Client.Collections().Create(schema); err != nil {
			log.Println("Error creating collection:", err)
			return
		}
	}

	// 1. Get a cursor for all users from MongoDB
	cursor, err := database.Database.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error fetching users:", err)
		return
	}
	defer cursor.Close(ctx)

	// Assuming `users` collection in Typesense
	collection := database.Client.Collection("users")

	// Helper function for upserting document
	upsertDocument := func(document map[string]interface{}, documentID string) {
		if _, err := collection.Documents().Upsert(document); err != nil {
			log.Printf("Error upserting document %s: %v\n", documentID, err)
		}
	}

	// Iterate through each user
	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Error decoding user:", err)
			continue
		}

		document := user.Document() // Assuming you have a Document method as before
		documentID := user.ID.Hex()

		// 2. Check if user exists in Typesense
		tsDocument, err := collection.Document(documentID).Retrieve()

		// If there's an error (like document doesn't exist) or the update timestamps differ, upsert
		if err != nil || tsDocument["updated_at"].(string) != document["updated_at"].(string) {
			upsertDocument(document, documentID)
		}
	}

	// Check if cursor encountered any error during iteration
	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
	}
}
