package service

import "github.com/google/wire"

var Set = wire.NewSet(
	NewSchedule,
	wire.Bind(new(Schedule), new(*ScheduleImpl)),
	NewUser,
	wire.Bind(new(User), new(*UserImpl)),
)
