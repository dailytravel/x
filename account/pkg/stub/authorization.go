package stub

import (
	"context"
	"log"
	"time"

	"github.com/dailytravel/x/account/pkg/database"
	"github.com/dailytravel/x/proto/account"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) Authorization(ctx context.Context, in *account.Request) (*account.Response, error) {
	// Convert the input message (token ID) to an ObjectID
	id, err := primitive.ObjectIDFromHex(in.Message)
	if err != nil {
		log.Printf("Failed to convert to ObjectID: %v", err)
		return nil, err
	}

	// Find the token by ID
	result := database.Database.Collection("tokens").FindOneAndUpdate(ctx,
		primitive.M{"_id": id, "revoked": false, "expires": bson.M{"$gt": primitive.Timestamp{T: uint32(time.Now().Unix())}}},
		bson.M{"$set": bson.M{"last_used": primitive.Timestamp{T: uint32(time.Now().Unix())}}})

	if result.Err() != nil {
		log.Printf("Failed to find and update token: %v", result.Err())
		return nil, result.Err()
	}

	var token map[string]interface{}
	if err := result.Decode(&token); err != nil {
		log.Printf("Failed to decode result: %v", err)
		return nil, err
	}

	// Find the user by user ID from the token
	err = database.Database.Collection("users").FindOne(ctx, primitive.M{"_id": token["uid"]}).Err()
	if err != nil {
		log.Printf("Failed to find user: %v", err)
		return nil, err
	}

	// Construct the response message
	response := &account.Response{
		Status: "authenticated",
	}

	return response, nil
}
