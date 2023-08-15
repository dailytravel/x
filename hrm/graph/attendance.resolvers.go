package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/hrm/graph/model"
	"github.com/dailytravel/x/hrm/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ID is the resolver for the id field.
func (r *attendanceResolver) ID(ctx context.Context, obj *model.Attendance) (string, error) {
	return obj.ID.Hex(), nil
}

// Employee is the resolver for the employee field.
func (r *attendanceResolver) Employee(ctx context.Context, obj *model.Attendance) (string, error) {
	panic(fmt.Errorf("not implemented: Employee - employee"))
}

// TimeIn is the resolver for the time_in field.
func (r *attendanceResolver) TimeIn(ctx context.Context, obj *model.Attendance) (string, error) {
	return time.Unix(int64(obj.TimeIn.T), 0).Format(time.RFC3339), nil
}

// TimeOut is the resolver for the time_out field.
func (r *attendanceResolver) TimeOut(ctx context.Context, obj *model.Attendance) (string, error) {
	return time.Unix(int64(obj.TimeOut.T), 0).Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *attendanceResolver) Metadata(ctx context.Context, obj *model.Attendance) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *attendanceResolver) CreatedAt(ctx context.Context, obj *model.Attendance) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *attendanceResolver) UpdatedAt(ctx context.Context, obj *model.Attendance) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *attendanceResolver) CreatedBy(ctx context.Context, obj *model.Attendance) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - created_by"))
}

// UpdatedBy is the resolver for the updated_by field.
func (r *attendanceResolver) UpdatedBy(ctx context.Context, obj *model.Attendance) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatedBy - updated_by"))
}

// CheckIn is the resolver for the checkIn field.
func (r *mutationResolver) CheckIn(ctx context.Context) (*model.Attendance, error) {
	// user := auth.User(ctx)
	// if user == nil {
	// 	return nil, fmt.Errorf("unauthorized")
	// }

	filter := bson.M{
		// "employee": user.ID,
		"time_in": bson.M{
			"$gte": primitive.Timestamp{T: uint32(time.Now().Unix()) - 86400},
			"$lt":  primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
		"time_out": bson.M{
			"$exists": false,
		},
	}

	var item *model.Attendance
	err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			item = &model.Attendance{
				// Employee: user.ID,
				TimeIn: primitive.Timestamp{T: uint32(time.Now().Unix())},
			}

			res, err := r.db.Collection(item.Collection()).InsertOne(ctx, item, nil)
			if err != nil {
				return nil, err
			}

			item.ID = res.InsertedID.(primitive.ObjectID)

			return item, nil
		}

		return nil, err
	}

	return item, nil
}

// CheckOut is the resolver for the checkOut field.
func (r *mutationResolver) CheckOut(ctx context.Context) (*model.Attendance, error) {
	// user := auth.User(ctx)
	// if user == nil {
	// 	return nil, fmt.Errorf("unauthorized")
	// }

	filter := bson.M{
		// "owner": user.ID,
		"time_in": bson.M{
			"$gte": primitive.Timestamp{T: uint32(time.Now().Unix()) - 86400},
			"$lt":  primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	var item model.Attendance

	err := r.db.Collection(item.Collection()).FindOne(ctx, filter, options.FindOne().SetSort(bson.D{{Key: "time_in", Value: -1}})).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("attendance not found")
		}
		return nil, err
	}

	out := primitive.Timestamp{T: uint32(time.Now().Unix())}
	update := &model.Attendance{
		TimeOut: &out,
	}

	_, err = r.db.Collection(item.Collection()).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// CreateAttendance is the resolver for the createAttendance field.
func (r *mutationResolver) CreateAttendance(ctx context.Context, input model.NewAttendance) (*model.Attendance, error) {
	panic(fmt.Errorf("not implemented: CreateAttendance - createAttendance"))
}

// UpdateAttendance is the resolver for the updateAttendance field.
func (r *mutationResolver) UpdateAttendance(ctx context.Context, id string, input model.UpdateAttendance) (*model.Attendance, error) {
	panic(fmt.Errorf("not implemented: UpdateAttendance - updateAttendance"))
}

// DeleteAttendance is the resolver for the deleteAttendance field.
func (r *mutationResolver) DeleteAttendance(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteAttendance - deleteAttendance"))
}

// DeleteAttendances is the resolver for the deleteAttendances field.
func (r *mutationResolver) DeleteAttendances(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteAttendances - deleteAttendances"))
}

// Attendance is the resolver for the attendance field.
func (r *queryResolver) Attendance(ctx context.Context, id string) (*model.Attendance, error) {
	panic(fmt.Errorf("not implemented: Attendance - attendance"))
}

// Attendances is the resolver for the attendances field.
func (r *queryResolver) Attendances(ctx context.Context, args map[string]interface{}) (*model.Attendances, error) {
	var items []*model.Attendance

	filter := utils.Query(args).(bson.M)

	opts := utils.Options(args)
	opts.SetSort(bson.M{"time_in": -1})

	cur, err := r.db.Collection("attendances").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Attendance
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("attendances").CountDocuments(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	return &model.Attendances{
		Count: int(count),
		Data:  items,
	}, nil
}

// Attendance returns AttendanceResolver implementation.
func (r *Resolver) Attendance() AttendanceResolver { return &attendanceResolver{r} }

type attendanceResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *attendanceResolver) Owner(ctx context.Context, obj *model.Attendance) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}
