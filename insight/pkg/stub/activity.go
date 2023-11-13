package stub

import (
	"context"
	"fmt"
	"log"

	"github.com/dailytravel/x/account/pkg/database"
	"github.com/dailytravel/x/proto/activity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) Activity(ctx context.Context, in *activity.Request) (*activity.Response, error) {
	uid, err := primitive.ObjectIDFromHex(in.Uid)
	if err != nil {
		log.Printf("Failed to convert to ObjectID: %v", err)
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(in.ObjectId)
	if err != nil {
		log.Printf("Failed to convert to ObjectID: %v", err)
		return nil, err
	}

	doc := bson.M{
		"uid":         uid,
		"action":      in.Action,
		"object._id":  objectId,
		"object.type": in.ObjectType,
		"timestamp":   primitive.Timestamp{T: uint32(in.Timestamp)},
	}

	if in.TargetId != "" {
		targetId, err := primitive.ObjectIDFromHex(in.TargetId)
		if err != nil {
			log.Printf("Failed to convert to ObjectID: %v", err)
			return nil, err
		}
		doc["target._id"] = targetId
		doc["target.type"] = in.TargetType
	}

	res, err := database.Database.Collection("activities").InsertOne(ctx, doc)
	if err != nil {
		log.Printf("Failed to insert activity: %v", err)
		return nil, err
	}

	response := &activity.Response{
		Status:  "success",
		Message: fmt.Sprintf("Inserted %v documents into activity collection", res.InsertedID),
	}

	return response, nil
}
