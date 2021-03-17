package resolver

import (
	"github.com/ww24/graphql-tutorial/domain/service"
	"github.com/ww24/graphql-tutorial/presentation/graphql/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	schedule service.Schedule
	user     service.User
}

func NewResolver(
	schedule service.Schedule,
	user service.User,
) *Resolver {
	return &Resolver{
		schedule: schedule,
		user:     user,
	}
}

func New(resolver *Resolver) generated.Config {
	c := generated.Config{
		Resolvers: resolver,
	}
	// TODO: Directives.HasRole
	// TODO: DIrectives.User
	return c
}
