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

type User interface {
	List(context.Context) ([]*model.User, error)
	GetByIDs(context.Context, []string) ([]*model.User, error)
	Create(context.Context, *model.NewUser) (*model.User, error)
}

type UserImpl struct {
	user repository.User
}

func NewUser(user repository.User) *UserImpl {
	return &UserImpl{user: user}
}

func (s *UserImpl) List(ctx context.Context) ([]*model.User, error) {
	log.Println("(*UserImpl).List called")
	return s.user.List(ctx)
}

func (s *UserImpl) GetByIDs(ctx context.Context, ids []string) ([]*model.User, error) {
	log.Println("(*UserImpl).GetByIDs called")
	return s.user.GetByIDs(ctx, ids)
}

func (s *UserImpl) Create(ctx context.Context, in *model.NewUser) (*model.User, error) {
	log.Println("(*UserImpl).Create called")
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("failed to generate UUIDv4")
	}

	user := &model.User{
		ID:   id.String(),
		Name: in.Name,
	}

	if err := s.user.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}
