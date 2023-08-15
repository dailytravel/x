package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/sales/graph/model"
)

// CreateProgram is the resolver for the createProgram field.
func (r *mutationResolver) CreateProgram(ctx context.Context, input model.NewProgram) (*model.Program, error) {
	panic(fmt.Errorf("not implemented: CreateProgram - createProgram"))
}

// UpdateProgram is the resolver for the updateProgram field.
func (r *mutationResolver) UpdateProgram(ctx context.Context, id string, input model.UpdateProgram) (*model.Program, error) {
	panic(fmt.Errorf("not implemented: UpdateProgram - updateProgram"))
}

// DeleteProgram is the resolver for the deleteProgram field.
func (r *mutationResolver) DeleteProgram(ctx context.Context, id string) (*model.Program, error) {
	panic(fmt.Errorf("not implemented: DeleteProgram - deleteProgram"))
}

// DeletePrograms is the resolver for the deletePrograms field.
func (r *mutationResolver) DeletePrograms(ctx context.Context, ids []string) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeletePrograms - deletePrograms"))
}

// ID is the resolver for the id field.
func (r *programResolver) ID(ctx context.Context, obj *model.Program) (string, error) {
	return obj.ID.Hex(), nil
}

// Type is the resolver for the type field.
func (r *programResolver) Type(ctx context.Context, obj *model.Program) (model.ProgramType, error) {
	panic(fmt.Errorf("not implemented: Type - type"))
}

// Name is the resolver for the name field.
func (r *programResolver) Name(ctx context.Context, obj *model.Program) (string, error) {
	panic(fmt.Errorf("not implemented: Name - name"))
}

// Description is the resolver for the description field.
func (r *programResolver) Description(ctx context.Context, obj *model.Program) (string, error) {
	panic(fmt.Errorf("not implemented: Description - description"))
}

// Metadata is the resolver for the metadata field.
func (r *programResolver) Metadata(ctx context.Context, obj *model.Program) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *programResolver) CreatedAt(ctx context.Context, obj *model.Program) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *programResolver) UpdatedAt(ctx context.Context, obj *model.Program) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *programResolver) CreatedBy(ctx context.Context, obj *model.Program) (*string, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - created_by"))
}

// UpdatedBy is the resolver for the updated_by field.
func (r *programResolver) UpdatedBy(ctx context.Context, obj *model.Program) (*string, error) {
	panic(fmt.Errorf("not implemented: UpdatedBy - updated_by"))
}

// Program is the resolver for the program field.
func (r *queryResolver) Program(ctx context.Context, id string) (*model.Program, error) {
	panic(fmt.Errorf("not implemented: Program - program"))
}

// Programs is the resolver for the programs field.
func (r *queryResolver) Programs(ctx context.Context, args map[string]interface{}) (*model.Programs, error) {
	panic(fmt.Errorf("not implemented: Programs - programs"))
}

// Program returns ProgramResolver implementation.
func (r *Resolver) Program() ProgramResolver { return &programResolver{r} }

type programResolver struct{ *Resolver }
