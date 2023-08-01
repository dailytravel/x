package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/account/graph/model"
)

// UpdateNotification is the resolver for the updateNotification field.
func (r *mutationResolver) UpdateNotification(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: UpdateNotification - updateNotification"))
}

// DeleteNotification is the resolver for the deleteNotification field.
func (r *mutationResolver) DeleteNotification(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteNotification - deleteNotification"))
}

// DeleteNotifications is the resolver for the deleteNotifications field.
func (r *mutationResolver) DeleteNotifications(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteNotifications - deleteNotifications"))
}

// Notifications is the resolver for the notifications field.
func (r *queryResolver) Notifications(ctx context.Context, args map[string]interface{}) (*model.Notifications, error) {
	panic(fmt.Errorf("not implemented: Notifications - notifications"))
}

// Notification is the resolver for the notification field.
func (r *queryResolver) Notification(ctx context.Context, id string) (*model.Notification, error) {
	panic(fmt.Errorf("not implemented: Notification - notification"))
}
