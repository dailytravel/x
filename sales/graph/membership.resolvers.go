package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/sales/graph/model"
	"github.com/dailytravel/x/sales/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ID is the resolver for the id field.
func (r *membershipResolver) ID(ctx context.Context, obj *model.Membership) (string, error) {
	return obj.ID.Hex(), nil
}

// Tier is the resolver for the tier field.
func (r *membershipResolver) Tier(ctx context.Context, obj *model.Membership) (*model.Tier, error) {
	var item *model.Tier

	filter := bson.M{"_id": obj.Tier}
	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("document not found")
		}
		return nil, err
	}

	return item, nil
}

// Billing is the resolver for the billing field.
func (r *membershipResolver) Billing(ctx context.Context, obj *model.Membership) (map[string]interface{}, error) {
	return obj.Billing, nil
}

// Payment is the resolver for the payment field.
func (r *membershipResolver) Payment(ctx context.Context, obj *model.Membership) (map[string]interface{}, error) {
	return obj.Payment, nil
}

// Transactions is the resolver for the transactions field.
func (r *membershipResolver) Transactions(ctx context.Context, obj *model.Membership) ([]*model.Transaction, error) {
	// Here you would typically query your database to fetch the transactions associated with the membership.
	// You might need to adjust the query and collection name based on your schema.

	var items []*model.Transaction

	// Example query using MongoDB
	filter := bson.M{"membership": obj.ID} // Assuming you have a field named membership_id in the Transaction schema
	cur, err := r.db.Collection("transactions").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var transaction model.Transaction
		if err := cur.Decode(&transaction); err != nil {
			return nil, err
		}
		items = append(items, &transaction)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// Metadata is the resolver for the metadata field.
func (r *membershipResolver) Metadata(ctx context.Context, obj *model.Membership) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *membershipResolver) CreatedAt(ctx context.Context, obj *model.Membership) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *membershipResolver) UpdatedAt(ctx context.Context, obj *model.Membership) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *membershipResolver) UID(ctx context.Context, obj *model.Membership) (string, error) {
	return obj.ID.Hex(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *membershipResolver) CreatedBy(ctx context.Context, obj *model.Membership) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *membershipResolver) UpdatedBy(ctx context.Context, obj *model.Membership) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// CreateMembership is the resolver for the createMembership field.
func (r *mutationResolver) CreateMembership(ctx context.Context, input model.NewMembership) (*model.Membership, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	tier, err := primitive.ObjectIDFromHex(input.Tier)
	if err != nil {
		return nil, err
	}

	item := &model.Membership{
		UID:     *uid,
		Tier:    tier,
		Number:  utils.GenerateMembershipCardNumber(),
		Since:   input.Since,
		Until:   input.Until,
		Billing: input.Billing,
		Payment: input.Payment,
		Model: model.Model{
			CreatedBy: uid,
			UpdatedBy: uid,
			Metadata:  input.Metadata,
		},
	}

	// Set the fields from the input
	_, err = r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateMembership is the resolver for the updateMembership field.
func (r *mutationResolver) UpdateMembership(ctx context.Context, id string, input model.UpdateMembership) (*model.Membership, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Create an update document with the fields to be updated
	item := &model.Membership{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	if input.Tier != nil {
		tier, err := primitive.ObjectIDFromHex(*input.Tier)
		if err != nil {
			return nil, err
		}
		item.Tier = tier
	}

	if input.Billing != nil {
		item.Billing = input.Billing
	}

	if input.Payment != nil {
		item.Payment = input.Payment
	}

	if input.Since != nil {
		item.Since = *input.Since
	}

	if input.Until != nil {
		item.Until = *input.Until
	}

	if input.Metadata != nil {
		for k, v := range input.Metadata {
			item.Metadata[k] = v
		}
	}

	if input.Status != nil {
		item.Status = *input.Status
	}

	// Update the updated_by and updated_at fields
	item.UpdatedBy = uid

	// Perform the update in the database
	res, err := r.db.Collection(item.Collection()).UpdateOne(ctx, filter, item)
	if err != nil {
		return nil, err
	}

	// Check if the coupon was actually updated
	if res.ModifiedCount == 0 {
		return nil, fmt.Errorf("no coupon was updated")
	}

	return item, nil
}

// DeleteMembership is the resolver for the deleteMembership field.
func (r *mutationResolver) DeleteMembership(ctx context.Context, id string) (*model.Membership, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the item by ID
	item := &model.Membership{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Mark the membership as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update operation in the database
	_, err = r.db.Collection(item.Collection()).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteMemberships is the resolver for the deleteMemberships field.
func (r *mutationResolver) DeleteMemberships(ctx context.Context, ids []string) ([]*model.Membership, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the list of ID strings to ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, _id)
	}

	// Define the filter to match the given IDs
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Define the update to mark records as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update operation in the database
	_, err = r.db.Collection("memberships").UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Fetch and return the deleted memberships
	var deletedMemberships []*model.Membership
	cursor, err := r.db.Collection("memberships").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		membership := &model.Membership{}
		if err := cursor.Decode(membership); err != nil {
			return nil, err
		}
		deletedMemberships = append(deletedMemberships, membership)
	}

	return deletedMemberships, nil
}

// Membership is the resolver for the membership field.
func (r *queryResolver) Membership(ctx context.Context, id string) (*model.Membership, error) {
	var item *model.Membership

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}
	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("document not found")
		}
		return nil, err
	}

	return item, nil
}

// Memberships is the resolver for the memberships field.
func (r *queryResolver) Memberships(ctx context.Context, args map[string]interface{}) (*model.Memberships, error) {
	var items []*model.Membership
	//find all items
	cur, err := r.db.Collection("memberships").Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Membership
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("memberships").CountDocuments(ctx, utils.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Memberships{
		Count: int(count),
		Data:  items,
	}, nil
}

// Membership returns MembershipResolver implementation.
func (r *Resolver) Membership() MembershipResolver { return &membershipResolver{r} }

type membershipResolver struct{ *Resolver }
