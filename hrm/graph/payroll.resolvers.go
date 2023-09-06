package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"time"

	"github.com/dailytravel/x/hrm/graph/model"
	"github.com/dailytravel/x/hrm/internal/utils"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreatePayroll is the resolver for the createPayroll field.
func (r *mutationResolver) CreatePayroll(ctx context.Context, input model.NewPayroll) (*model.Payroll, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	item := &model.Payroll{
		Amount:   input.Amount,
		Currency: input.Currency,
		Status:   input.Status,
		Model: model.Model{
			CreatedBy: uid,
			UpdatedBy: uid,
			Metadata:  input.Metadata,
		},
	}

	// Convert the ID string to ObjectID
	_uid, err := primitive.ObjectIDFromHex(input.UID)
	if err != nil {
		return nil, err
	}

	item.UID = _uid

	//convert date string to primitive.DateTime

	payDate, err := time.Parse(time.RFC3339, input.PayDate)
	if err != nil {
		return nil, err
	}

	item.PayDate = primitive.NewDateTimeFromTime(payDate)

	_, err = r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdatePayroll is the resolver for the updatePayroll field.
func (r *mutationResolver) UpdatePayroll(ctx context.Context, id string, input model.UpdatePayroll) (*model.Payroll, error) {
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
	item := &model.Payroll{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	if input.Amount != nil {
		item.Amount = *input.Amount
	}

	if input.Currency != nil {
		item.Currency = *input.Currency
	}

	if input.Status != nil {
		item.Status = *input.Status
	}

	if input.Metadata != nil {
		for k, v := range input.Metadata {
			item.Metadata[k] = v
		}
	}

	if input.PayDate != nil {
		payDate, err := time.Parse(time.RFC3339, *input.PayDate)
		if err != nil {
			return nil, err
		}

		item.PayDate = primitive.NewDateTimeFromTime(payDate)
	}

	item.UpdatedBy = uid

	if err := r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, item).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

// DeletePayroll is the resolver for the deletePayroll field.
func (r *mutationResolver) DeletePayroll(ctx context.Context, id string) (*model.Payroll, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the payroll record by ID
	item := &model.Payroll{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update the payroll record to mark it as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Perform the update operation in the database
	var deletedPayroll model.Payroll
	err = r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, update, opts).Decode(&deletedPayroll)
	if err != nil {
		return nil, err
	}

	return &deletedPayroll, nil
}

// DeletePayrolls is the resolver for the deletePayrolls field.
func (r *mutationResolver) DeletePayrolls(ctx context.Context, ids []string) ([]*model.Payroll, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	var deletedPayrolls []*model.Payroll

	for _, id := range ids {
		// Convert the ID string to ObjectID
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}

		// Find the payroll record by ID
		item := &model.Payroll{}
		filter := bson.M{"_id": _id}
		err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
		if err != nil {
			return nil, err
		}

		// Update the payroll record to mark it as deleted
		update := bson.M{
			"$set": bson.M{
				"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
				"deleted_by": uid,
				"status":     "deleted",
				"updated_by": uid,
				"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			},
		}

		opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

		// Perform the update operation in the database
		var deletedPayroll model.Payroll
		err = r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, update, opts).Decode(&deletedPayroll)
		if err != nil {
			return nil, err
		}

		deletedPayrolls = append(deletedPayrolls, &deletedPayroll)
	}

	return deletedPayrolls, nil
}

// ID is the resolver for the id field.
func (r *payrollResolver) ID(ctx context.Context, obj *model.Payroll) (string, error) {
	return obj.ID.Hex(), nil
}

// PayDate is the resolver for the pay_date field.
func (r *payrollResolver) PayDate(ctx context.Context, obj *model.Payroll) (string, error) {
	return obj.PayDate.Time().String(), nil
}

// Metadata is the resolver for the metadata field.
func (r *payrollResolver) Metadata(ctx context.Context, obj *model.Payroll) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *payrollResolver) CreatedAt(ctx context.Context, obj *model.Payroll) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *payrollResolver) UpdatedAt(ctx context.Context, obj *model.Payroll) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *payrollResolver) UID(ctx context.Context, obj *model.Payroll) (string, error) {
	return obj.UID.Hex(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *payrollResolver) CreatedBy(ctx context.Context, obj *model.Payroll) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	return pointer.String(obj.CreatedBy.Hex()), nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *payrollResolver) UpdatedBy(ctx context.Context, obj *model.Payroll) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	return pointer.String(obj.UpdatedBy.Hex()), nil
}

// Payroll is the resolver for the payroll field.
func (r *queryResolver) Payroll(ctx context.Context, id string) (*model.Payroll, error) {
	var item *model.Payroll

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}

	if err := r.db.Collection("payrolls").FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// Payrolls is the resolver for the payrolls field.
func (r *queryResolver) Payrolls(ctx context.Context, args map[string]interface{}) (*model.Payrolls, error) {
	var items []*model.Payroll
	//find all items
	cur, err := r.db.Collection("payrolls").Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Payroll
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("payrolls").CountDocuments(ctx, utils.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Payrolls{
		Count: int(count),
		Data:  items,
	}, nil
}

// Payroll returns PayrollResolver implementation.
func (r *Resolver) Payroll() PayrollResolver { return &payrollResolver{r} }

type payrollResolver struct{ *Resolver }
