package repository

import (
	"context"

	"github.com/ww24/graphql-tutorial/presentation/graphql/model"
)

type User interface {
	GetByIDs(context.Context, []string) ([]*model.User, error)
	List(context.Context) ([]*model.User, error)
	Save(context.Context, *model.User) error
}
