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
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ID is the resolver for the id field.
func (r *companyResolver) ID(ctx context.Context, obj *model.Company) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *companyResolver) Metadata(ctx context.Context, obj *model.Company) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *companyResolver) CreatedAt(ctx context.Context, obj *model.Company) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *companyResolver) UpdatedAt(ctx context.Context, obj *model.Company) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *companyResolver) UID(ctx context.Context, obj *model.Company) (string, error) {
	return obj.ID.Hex(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *companyResolver) CreatedBy(ctx context.Context, obj *model.Company) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	return pointer.String(obj.CreatedBy.Hex()), nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *companyResolver) UpdatedBy(ctx context.Context, obj *model.Company) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	return pointer.String(obj.UpdatedBy.Hex()), nil
}

// CreateCompany is the resolver for the createCompany field.
func (r *mutationResolver) CreateCompany(ctx context.Context, input model.NewCompany) (*model.Company, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	item := &model.Company{
		Type:        input.Type,
		Name:        input.Name,
		Description: input.Description,
		Industry:    input.Industry,
		Employees:   input.Employees,
		Revenue:     input.Revenue,
		City:        input.City,
		Zip:         input.Zip,
		State:       input.State,
		Country:     input.Country,
		Timezone:    input.Timezone,
		Phone:       input.Phone,
		Website:     input.Website,
		Status:      input.Status,
		Model: model.Model{
			Metadata:  input.Metadata,
			CreatedBy: uid,
			UpdatedBy: uid,
		},
	}

	// Set the fields from the input
	_, err = r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateCompany is the resolver for the updateCompany field.
func (r *mutationResolver) UpdateCompany(ctx context.Context, id string, input model.UpdateCompany) (*model.Company, error) {
	// Get the authenticated user ID
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Retrieve the existing company from the database
	item := &model.Company{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update the company fields based on the input data
	// For example:
	if input.Name != nil {
		item.Name = *input.Name
	}

	if input.Description != nil {
		item.Description = input.Description
	}

	if input.Industry != nil {
		item.Industry = input.Industry
	}

	if input.Employees != nil {
		item.Employees = input.Employees
	}

	if input.Revenue != nil {
		item.Revenue = input.Revenue
	}

	if input.City != nil {
		item.City = input.City
	}

	if input.Zip != nil {
		item.Zip = input.Zip
	}

	if input.State != nil {
		item.State = input.State
	}

	if input.Country != nil {
		item.Country = input.Country
	}

	if input.Timezone != nil {
		item.Timezone = input.Timezone
	}

	if input.Phone != nil {
		item.Phone = input.Phone
	}

	if input.Website != nil {
		item.Website = input.Website
	}

	if input.Status != nil {
		item.Status = *input.Status
	}

	// Update the "updated_by" field with the user ID
	item.UpdatedBy = uid

	// Perform the update operation in the database
	updatedCompany := &model.Company{}
	updateResult := r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, bson.M{"$set": item})
	if updateResult.Err() != nil {
		return nil, updateResult.Err()
	}

	// Decode the updated company
	err = updateResult.Decode(updatedCompany)
	if err != nil {
		return nil, err
	}

	return updatedCompany, nil
}

// DeleteCompany is the resolver for the deleteCompany field.
func (r *mutationResolver) DeleteCompany(ctx context.Context, id string) (map[string]interface{}, error) {
	// Convert the ID string to an ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Define the filter to match the given ID
	filter := bson.M{"_id": _id}

	// Delete the company from the database
	result, err := r.db.Collection("companies").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("company not found")
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.DeletedCount}, nil
}

// DeleteCompanies is the resolver for the deleteCompanies field.
func (r *mutationResolver) DeleteCompanies(ctx context.Context, ids []*string) (map[string]interface{}, error) {
	// Convert the list of ID strings to ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(*id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, _id)
	}

	// Define the filter to match the given IDs
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Delete the companies from the database
	result, err := r.db.Collection("companies").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.DeletedCount}, nil
}

// Companies is the resolver for the companies field.
func (r *queryResolver) Companies(ctx context.Context, args map[string]interface{}) (*model.Companies, error) {
	var items []*model.Company
	//find all items
	cur, err := r.db.Collection("companies").Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Company
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("companies").CountDocuments(ctx, utils.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Companies{
		Count: int(count),
		Data:  items,
	}, nil
}

// Company is the resolver for the company field.
func (r *queryResolver) Company(ctx context.Context, id string) (*model.Company, error) {
	var item *model.Company

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

// Company returns CompanyResolver implementation.
func (r *Resolver) Company() CompanyResolver { return &companyResolver{r} }

type companyResolver struct{ *Resolver }
