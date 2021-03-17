package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/ww24/graphql-tutorial/presentation/graphql/generated"
	"github.com/ww24/graphql-tutorial/presentation/graphql/model"
)

func (r *queryResolver) Schedule(ctx context.Context, id string) (*model.Schedule, error) {
	schedules, err := r.schedule.GetByIDs(ctx, []string{id})
	if err != nil {
		return nil, fmt.Errorf("failed to get schedule: %w", err)
	}
	if len(schedules) == 0 {
		return nil, errors.New("schedule not found")
	}
	return schedules[0], nil
}

func (r *queryResolver) Schedules(ctx context.Context, ids []string) ([]*model.Schedule, error) {
	return r.schedule.GetByIDs(ctx, ids)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	users, err := r.user.GetByIDs(ctx, []string{id})
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if len(users) == 0 {
		return nil, errors.New("user not found")
	}
	return users[0], nil
}

func (r *queryResolver) Users(ctx context.Context, ids []string) ([]*model.User, error) {
	return r.user.GetByIDs(ctx, ids)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
