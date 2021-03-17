package db

import (
	"context"
	"sync"

	"github.com/ww24/graphql-tutorial/presentation/graphql/model"
)

type Schedule struct {
	mu        sync.Mutex
	schedules []*model.Schedule
	index     map[string]*model.Schedule
}

func NewSchedule() *Schedule {
	return &Schedule{
		index: make(map[string]*model.Schedule),
	}
}

func (s *Schedule) GetByIDs(ctx context.Context, ids []string) ([]*model.Schedule, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	schedules := make([]*model.Schedule, 0, len(ids))
	for _, id := range ids {
		if schedule, ok := s.index[id]; ok {
			schedules = append(schedules, schedule)
		}
	}
	return schedules, nil
}

func (s *Schedule) GetByUserIDs(ctx context.Context, userIDs []string) ([]*model.Schedule, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userMap := make(map[string]struct{}, len(userIDs))
	for _, id := range userIDs {
		userMap[id] = struct{}{}
	}

	schedules := make([]*model.Schedule, 0)
	for _, schedule := range s.schedules {
		if schedule.CreatedBy == nil {
			continue
		}
		if _, ok := userMap[schedule.CreatedBy.ID]; ok {
			schedules = append(schedules, schedule)
		}
	}
	return schedules, nil
}

func (s *Schedule) List(ctx context.Context) ([]*model.Schedule, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.schedules, nil
}

func (s *Schedule) Save(ctx context.Context, schedule *model.Schedule) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.schedules = append(s.schedules, schedule)
	s.index[schedule.ID] = schedule
	return nil
}
