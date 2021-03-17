package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/ww24/graphql-tutorial/domain/repository"
	"github.com/ww24/graphql-tutorial/presentation/graphql/model"
)

type Schedule interface {
	List(context.Context) ([]*model.Schedule, error)
	GetByIDs(context.Context, []string) ([]*model.Schedule, error)
	GetByUserIDs(ctx context.Context, ids []string) ([]*model.Schedule, error)
	Create(context.Context, *model.NewSchedule) (*model.Schedule, error)
}

type ScheduleImpl struct {
	schedule repository.Schedule
	user     repository.User
}

func NewSchedule(schedule repository.Schedule, user repository.User) *ScheduleImpl {
	return &ScheduleImpl{
		schedule: schedule,
		user:     user,
	}
}

func (s *ScheduleImpl) List(ctx context.Context) ([]*model.Schedule, error) {
	log.Println("(*ScheduleImpl).List called")
	return s.schedule.List(ctx)
}

func (s *ScheduleImpl) GetByIDs(ctx context.Context, ids []string) ([]*model.Schedule, error) {
	log.Println("(*ScheduleImpl).GetByIDs called")
	return s.schedule.GetByIDs(ctx, ids)
}

func (s *ScheduleImpl) GetByUserIDs(ctx context.Context, ids []string) ([]*model.Schedule, error) {
	log.Println("(*ScheduleImpl).GetByUserIDs called")
	return s.schedule.GetByUserIDs(ctx, ids)
}

func (s *ScheduleImpl) Create(ctx context.Context, in *model.NewSchedule) (*model.Schedule, error) {
	log.Println("(*ScheduleImpl).Create called")
	participants, err := s.user.GetByIDs(ctx, in.ParticipantsUserIds)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	users, err := s.user.GetByIDs(ctx, []string{in.UserID})
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if len(users) == 0 {
		return nil, errors.New("user not found")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("failed to generate UUIDv4")
	}
	schedule := &model.Schedule{
		ID:           id.String(),
		Title:        in.Title,
		StartAt:      in.StartAt,
		EndAt:        in.EndAt,
		CreatedBy:    users[0],
		Description:  in.Description,
		Participants: participants,
	}

	if err := s.schedule.Save(ctx, schedule); err != nil {
		return nil, fmt.Errorf("failed to save schedule: %w", err)
	}
	return schedule, nil
}
