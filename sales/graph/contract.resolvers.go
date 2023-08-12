package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/sales/graph/model"
)

// ID is the resolver for the id field.
func (r *contractResolver) ID(ctx context.Context, obj *model.Contract) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// User is the resolver for the user field.
func (r *contractResolver) User(ctx context.Context, obj *model.Contract) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Contact is the resolver for the contact field.
func (r *contractResolver) Contact(ctx context.Context, obj *model.Contract) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: Contact - contact"))
}

// StartDate is the resolver for the start_date field.
func (r *contractResolver) StartDate(ctx context.Context, obj *model.Contract) (string, error) {
	panic(fmt.Errorf("not implemented: StartDate - start_date"))
}

// EndDate is the resolver for the end_date field.
func (r *contractResolver) EndDate(ctx context.Context, obj *model.Contract) (string, error) {
	panic(fmt.Errorf("not implemented: EndDate - end_date"))
}

// Metadata is the resolver for the metadata field.
func (r *contractResolver) Metadata(ctx context.Context, obj *model.Contract) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: Metadata - metadata"))
}

// CreatedAt is the resolver for the created_at field.
func (r *contractResolver) CreatedAt(ctx context.Context, obj *model.Contract) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - created_at"))
}

// UpdatedAt is the resolver for the updated_at field.
func (r *contractResolver) UpdatedAt(ctx context.Context, obj *model.Contract) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updated_at"))
}

// Comments is the resolver for the comments field.
func (r *contractResolver) Comments(ctx context.Context, obj *model.Contract) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// Followers is the resolver for the followers field.
func (r *contractResolver) Followers(ctx context.Context, obj *model.Contract) ([]*model.Follow, error) {
	panic(fmt.Errorf("not implemented: Followers - followers"))
}

// CreateContract is the resolver for the createContract field.
func (r *mutationResolver) CreateContract(ctx context.Context, input model.NewContract) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: CreateContract - createContract"))
}

// UpdateContract is the resolver for the updateContract field.
func (r *mutationResolver) UpdateContract(ctx context.Context, id string, input model.UpdateContract) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: UpdateContract - updateContract"))
}

// DeleteContract is the resolver for the deleteContract field.
func (r *mutationResolver) DeleteContract(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteContract - deleteContract"))
}

// DeleteContracts is the resolver for the deleteContracts field.
func (r *mutationResolver) DeleteContracts(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteContracts - deleteContracts"))
}

// Contract is the resolver for the contract field.
func (r *queryResolver) Contract(ctx context.Context, id string) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented: Contract - contract"))
}

// Contracts is the resolver for the contracts field.
func (r *queryResolver) Contracts(ctx context.Context, args map[string]interface{}) (*model.Contracts, error) {
	panic(fmt.Errorf("not implemented: Contracts - contracts"))
}

// Contract returns ContractResolver implementation.
func (r *Resolver) Contract() ContractResolver { return &contractResolver{r} }

type contractResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *contractResolver) Owner(ctx context.Context, obj *model.Contract) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}
