package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ww24/graphql-tutorial/presentation/graphql/generated"
	"github.com/ww24/graphql-tutorial/presentation/graphql/model"
)

func (r *mutationResolver) CreateSchedule(ctx context.Context, input model.NewSchedule) (*model.Schedule, error) {
	return r.schedule.Create(ctx, &input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.user.Create(ctx, &input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
