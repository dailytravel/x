package controllers

import (
	"context"
	"log"
	"sync"

	"github.com/dailytravel/x/cms/graph/model"
	"github.com/dailytravel/x/cms/pkg/database"
	"github.com/typesense/typesense-go/typesense"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Indexer interface {
	Collection() string
	Insert(collection typesense.CollectionInterface) error
	Update(collection typesense.CollectionInterface, documentKey primitive.M, updatedFields primitive.M, removedFields primitive.A) error
	Delete(collection typesense.CollectionInterface, documentKey primitive.M) error
}

func Index(item Indexer, data primitive.M) error {
	operationType := data["operationType"].(string)
	collection := database.Client.Collection(item.Collection())

	switch operationType {
	case "insert":
		fullDocument := data["fullDocument"].(primitive.M)

		// Convert fullDocument to bytes
		bytes, err := bson.Marshal(fullDocument)
		if err != nil {
			log.Printf("Error marshalling fullDocument: %v", err)
			return err
		}

		// Unmarshal bytes to the item struct
		err = bson.Unmarshal(bytes, item)
		if err != nil {
			log.Printf("Error unmarshalling fullDocument: %v", err)
			return err
		}

		// Code to handle insert/update operation
		err = item.Insert(collection)
		if err != nil {
			return err
		}

	case "update":
		documentKey := data["documentKey"].(primitive.M)
		updatedFields := data["updateDescription"].(primitive.M)["updatedFields"].(primitive.M)
		removedFields := data["updateDescription"].(primitive.M)["removedFields"].(primitive.A)

		if err := database.Database.Collection(item.Collection()).FindOne(context.TODO(), bson.M{"_id": documentKey["_id"].(primitive.ObjectID)}).Decode(item); err != nil {
			return err
		}

		err := item.Update(collection, documentKey, updatedFields, removedFields)
		if err != nil {
			return err
		}

	case "delete":
		documentKey := data["documentKey"].(primitive.M)

		err := item.Delete(collection, documentKey)
		if err != nil {
			return err
		}

	case "drop", "invalidate":
		if _, err := collection.Delete(); err != nil {
			return err
		}
	}

	return nil
}

func IndexStream(waitGroup *sync.WaitGroup, stream *mongo.ChangeStream, collectionName string) {
	defer stream.Close(context.TODO())

	indexers := map[string]Indexer{
		"terms":    &model.Term{},
		"contents": &model.Content{},
		"files":    &model.File{},
	}

	indexer, found := indexers[collectionName]
	if !found {
		log.Printf("Unknown collection name: %s", collectionName)
		waitGroup.Done() // Decrement the WaitGroup counter
		return
	}

	for stream.Next(context.TODO()) {
		var data primitive.M
		if err := stream.Decode(&data); err != nil {
			log.Printf("Error decoding change stream data for collection %s: %v", collectionName, err)
			continue
		}

		if err := Index(indexer, data); err != nil {
			log.Printf("Error indexing data from collection %s: %v", collectionName, err)
			continue
		}
	}

	if err := stream.Err(); err != nil {
		log.Printf("Error iterating change stream for collection %s: %v", collectionName, err)
	}

	waitGroup.Done() // Decrement the WaitGroup counter
}
