//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ww24/graphql-tutorial/domain/service"
	"github.com/ww24/graphql-tutorial/infra/db"
	"github.com/ww24/graphql-tutorial/presentation/graphql/resolver"
)

func register() *server {
	wire.Build(
		service.Set,
		db.Set,
		resolver.NewResolver,
		newServer,
	)
	return nil
}
