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
	"github.com/dailytravel/x/sales/internal/utils"
	"github.com/dailytravel/x/sales/pkg/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateVoucher is the resolver for the createVoucher field.
func (r *mutationResolver) CreateVoucher(ctx context.Context, input model.NewVoucher) (*model.Voucher, error) {
	// uid, err := utils.UID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	item := &model.Voucher{
		Locale: input.Locale,
		Description: bson.M{
			input.Locale: input.Description,
		},
		Model: model.Model{
			Metadata: input.Metadata,
		},
	}

	utils.Date(&input.Start, &item.Start)
	utils.Date(&input.End, &item.End)

	// Set the fields from the input
	_, err := r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateVoucher is the resolver for the updateVoucher field.
func (r *mutationResolver) UpdateVoucher(ctx context.Context, id string, input model.UpdateVoucher) (*model.Voucher, error) {
	// uid, err := utils.UID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Create an update document with the fields to be updated
	item := &model.Voucher{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	if input.Description != nil {
		item.Description[*input.Locale] = *input.Description
	}

	if input.Status != nil {
		item.Status = *input.Status
	}

	if input.Metadata != nil {
		for k, v := range input.Metadata {
			item.Metadata[k] = v
		}
	}

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

// DeleteVoucher is the resolver for the deleteVoucher field.
func (r *mutationResolver) DeleteVoucher(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Define the filter to match the given voucher ID
	filter := bson.M{"_id": objectID}

	// Define the update to mark the voucher as deleted
	update := bson.M{
		"$set": bson.M{
			"status":     "deleted",
			"updated_by": uid,
			"updated":    primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update operation in the database
	result, err := r.db.Collection("vouchers").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}, nil
}

// DeleteVouchers is the resolver for the deleteVouchers field.
func (r *mutationResolver) DeleteVouchers(ctx context.Context, ids []string) (map[string]interface{}, error) {
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

	// Define the filter to match the given voucher IDs
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Define the update to mark vouchers as deleted
	update := bson.M{
		"$set": bson.M{
			"status":     "deleted",
			"updated_by": uid,
			"updated":    primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update operation in the database
	result, err := r.db.Collection("vouchers").UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}, nil
}

// Voucher is the resolver for the voucher field.
func (r *queryResolver) Voucher(ctx context.Context, id string) (*model.Voucher, error) {
	var item *model.Voucher

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

// Vouchers is the resolver for the vouchers field.
func (r *queryResolver) Vouchers(ctx context.Context, stages map[string]interface{}) (*model.Vouchers, error) {
	var items []*model.Voucher
	//find all items
	cur, err := r.db.Collection("vouchers").Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Voucher
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("vouchers").CountDocuments(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &model.Vouchers{
		Count: int(count),
		Data:  items,
	}, nil
}

// ID is the resolver for the id field.
func (r *voucherResolver) ID(ctx context.Context, obj *model.Voucher) (string, error) {
	return obj.ID.Hex(), nil
}

// Name is the resolver for the name field.
func (r *voucherResolver) Name(ctx context.Context, obj *model.Voucher) (string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the name for the requested locale
	if name, ok := obj.Name[*locale].(string); ok {
		return name, nil
	}

	// Return an error if the name is not found for any locale
	return "", errors.New("Name not found for any locale")
}

// Description is the resolver for the description field.
func (r *voucherResolver) Description(ctx context.Context, obj *model.Voucher) (string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	// Try to retrieve the description for the requested locale
	if description, ok := obj.Description[*locale].(string); ok {
		return description, nil
	}

	// Return an error if the name is not found for any locale
	return "", errors.New("Description not found for any locale")
}

// Start is the resolver for the start field.
func (r *voucherResolver) Start(ctx context.Context, obj *model.Voucher) (string, error) {
	return time.Unix(int64(obj.Start), 0).Format(time.RFC3339), nil
}

// End is the resolver for the end field.
func (r *voucherResolver) End(ctx context.Context, obj *model.Voucher) (string, error) {
	return time.Unix(int64(obj.End), 0).Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *voucherResolver) Metadata(ctx context.Context, obj *model.Voucher) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Created is the resolver for the created field.
func (r *voucherResolver) Created(ctx context.Context, obj *model.Voucher) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *voucherResolver) Updated(ctx context.Context, obj *model.Voucher) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// Package is the resolver for the package field.
func (r *voucherResolver) Package(ctx context.Context, obj *model.Voucher) (string, error) {
	panic(fmt.Errorf("not implemented: Package - package"))
}

// Supplier is the resolver for the supplier field.
func (r *voucherResolver) Supplier(ctx context.Context, obj *model.Voucher) (string, error) {
	panic(fmt.Errorf("not implemented: Supplier - supplier"))
}

// Voucher returns VoucherResolver implementation.
func (r *Resolver) Voucher() VoucherResolver { return &voucherResolver{r} }

type voucherResolver struct{ *Resolver }
