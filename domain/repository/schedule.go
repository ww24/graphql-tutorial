package repository

import (
	"context"

	"github.com/ww24/graphql-tutorial/presentation/graphql/model"
)

type Schedule interface {
	GetByIDs(context.Context, []string) ([]*model.Schedule, error)
	GetByUserIDs(context.Context, []string) ([]*model.Schedule, error)
	List(context.Context) ([]*model.Schedule, error)
	Save(context.Context, *model.Schedule) error
}
