package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	dataloader "github.com/graph-gophers/dataloader/v6"
	dl "github.com/ww24/graphql-tutorial/presentation/graphql/dataloader"
	"github.com/ww24/graphql-tutorial/presentation/graphql/generated"
	"github.com/ww24/graphql-tutorial/presentation/graphql/model"
)

func (r *scheduleResolver) CreatedBy(ctx context.Context, obj *model.Schedule) (*model.User, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	// users, err := r.user.GetByIDs(ctx, []string{obj.CreatedBy.ID})
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get user: %w", err)
	// }
	// if len(users) == 0 {
	// 	return nil, errors.New("user not found")
	// }
	// return users[0], nil

	// use dataloader
	thunk := dl.FromContext(ctx).UserByID.Load(ctx, dataloader.StringKey(obj.CreatedBy.ID))
	data, err := thunk()
	if err != nil {
		return nil, fmt.Errorf("failed to get user by dataloader: %w", err)
	}
	return data.(*model.User), nil
}

// Schedule returns generated.ScheduleResolver implementation.
func (r *Resolver) Schedule() generated.ScheduleResolver { return &scheduleResolver{r} }

type scheduleResolver struct{ *Resolver }
