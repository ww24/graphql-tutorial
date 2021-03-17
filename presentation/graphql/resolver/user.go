package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ww24/graphql-tutorial/presentation/graphql/generated"
	"github.com/ww24/graphql-tutorial/presentation/graphql/model"
)

func (r *userResolver) Schedules(ctx context.Context, obj *model.User) ([]*model.Schedule, error) {
	return r.schedule.GetByUserIDs(ctx, []string{obj.ID})
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
