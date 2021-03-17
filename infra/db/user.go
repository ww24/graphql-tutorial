package db

import (
	"context"
	"sync"

	"github.com/ww24/graphql-tutorial/presentation/graphql/model"
)

type User struct {
	mu    sync.Mutex
	users []*model.User
	index map[string]*model.User
}

func NewUser() *User {
	return &User{
		index: make(map[string]*model.User),
	}
}

func (u *User) GetByIDs(ctx context.Context, ids []string) ([]*model.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	users := make([]*model.User, 0, len(ids))
	for _, id := range ids {
		if user, ok := u.index[id]; ok {
			users = append(users, user)
		}
	}
	return users, nil
}

func (u *User) List(ctx context.Context) ([]*model.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.users, nil
}

func (u *User) Save(ctx context.Context, user *model.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.users = append(u.users, user)
	u.index[user.ID] = user
	return nil
}
