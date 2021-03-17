package db

import (
	"github.com/google/wire"
	"github.com/ww24/graphql-tutorial/domain/repository"
)

var Set = wire.NewSet(
	NewSchedule,
	wire.Bind(new(repository.Schedule), new(*Schedule)),
	NewUser,
	wire.Bind(new(repository.User), new(*User)),
)
