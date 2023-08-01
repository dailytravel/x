package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/service/graph/model"
)

// CreateGoal is the resolver for the createGoal field.
func (r *mutationResolver) CreateGoal(ctx context.Context, input model.NewGoal) (*model.Goal, error) {
	panic(fmt.Errorf("not implemented: CreateGoal - createGoal"))
}

// UpdateGoal is the resolver for the updateGoal field.
func (r *mutationResolver) UpdateGoal(ctx context.Context, id string, input model.UpdateGoal) (*model.Goal, error) {
	panic(fmt.Errorf("not implemented: UpdateGoal - updateGoal"))
}

// DeleteGoal is the resolver for the deleteGoal field.
func (r *mutationResolver) DeleteGoal(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteGoal - deleteGoal"))
}

// DeleteGoals is the resolver for the deleteGoals field.
func (r *mutationResolver) DeleteGoals(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteGoals - deleteGoals"))
}

// CreateGoalMetric is the resolver for the createGoalMetric field.
func (r *mutationResolver) CreateGoalMetric(ctx context.Context, id string, input model.NewMetric) (*model.Goal, error) {
	panic(fmt.Errorf("not implemented: CreateGoalMetric - createGoalMetric"))
}

// UpdateGoalMetric is the resolver for the updateGoalMetric field.
func (r *mutationResolver) UpdateGoalMetric(ctx context.Context, id string, input model.UpdateMetric) (*model.Goal, error) {
	panic(fmt.Errorf("not implemented: UpdateGoalMetric - updateGoalMetric"))
}

// Goal is the resolver for the goal field.
func (r *queryResolver) Goal(ctx context.Context, id string) (*model.Goal, error) {
	panic(fmt.Errorf("not implemented: Goal - goal"))
}

// Goals is the resolver for the goals field.
func (r *queryResolver) Goals(ctx context.Context, args map[string]interface{}) (*model.Goals, error) {
	panic(fmt.Errorf("not implemented: Goals - goals"))
}
