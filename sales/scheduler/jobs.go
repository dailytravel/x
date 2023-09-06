package scheduler

import (
	"context"
	"log"

	"github.com/dailytravel/x/sales/graph/model"
	"github.com/dailytravel/x/sales/pkg/database"
	"github.com/robfig/cron"
	"github.com/typesense/typesense-go/typesense/api"
	"go.mongodb.org/mongo-driver/bson"
)

func SyncContactsJob() {
	c := cron.New()
	// c.AddFunc("@daily", func() {
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Starting synchronization...")
		syncContactsToTypesense()
		log.Println("Synchronization complete.")
	})
	c.Start()
}

func syncContactsToTypesense() {
	ctx := context.Background()

	contactModel := &model.Contact{}
	if _, err := database.Client.Collection(contactModel.Collection()).Retrieve(); err != nil {
		// Assuming a 'Schema' method on model.Contact returns the required collection schema for Typesense
		schema, ok := contactModel.Schema().(*api.CollectionSchema)
		if !ok {
			log.Println("Error casting schema to *api.CollectionSchema")
			return
		}
		if _, err := database.Client.Collections().Create(schema); err != nil {
			log.Println("Error creating collection:", err)
			return
		}
	}

	// 1. Get a cursor for all contacts from MongoDB
	cursor, err := database.Database.Collection("contacts").Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error fetching contacts:", err)
		return
	}
	defer cursor.Close(ctx)

	// Assuming `contacts` collection in Typesense
	collection := database.Client.Collection("contacts")

	// Helper function for upserting document
	upsertDocument := func(document map[string]interface{}, documentID string) {
		if _, err := collection.Documents().Upsert(document); err != nil {
			log.Printf("Error upserting document %s: %v\n", documentID, err)
		}
	}

	// Iterate through each contact
	for cursor.Next(ctx) {
		var contact model.Contact
		if err := cursor.Decode(&contact); err != nil {
			log.Println("Error decoding contact:", err)
			continue
		}

		document := contact.Document() // Assuming you have a Document method as before
		documentID := contact.ID.Hex()

		// 2. Check if contact exists in Typesense
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
