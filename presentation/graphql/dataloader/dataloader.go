package dataloader

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/graph-gophers/dataloader/v6"
	"github.com/ww24/graphql-tutorial/domain/service"
	"github.com/ww24/graphql-tutorial/presentation/graphql/model"
)

var (
	key             = struct{}{}
	errUserNotFound = errors.New("user not found")
)

type Loader struct {
	user     service.User
	UserByID dataloader.Interface
}

func newLoader(user service.User) *Loader {
	l := &Loader{
		user: user,
	}
	l.UserByID = dataloader.NewBatchedLoader(l.userBatchFn, dataloader.WithWait(2*time.Millisecond))
	return l
}

func Middleware(user service.User) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := withLoader(r.Context(), newLoader(user))
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func withLoader(ctx context.Context, loader *Loader) context.Context {
	return context.WithValue(ctx, key, loader)
}

func (l *Loader) userBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	results := make([]*dataloader.Result, 0, len(keys))
	ids := make([]string, 0, len(keys))
	for _, k := range keys {
		ids = append(ids, k.String())
	}

	users, err := l.user.GetByIDs(ctx, ids)
	if err != nil {
		for range keys {
			results = append(results, &dataloader.Result{Error: err})
		}
		return results
	}

	userMap := make(map[string]*model.User, len(users))
	for _, user := range users {
		userMap[user.ID] = user
	}

	for _, id := range ids {
		result := &dataloader.Result{}
		if user, ok := userMap[id]; ok {
			result.Data = user
		} else {
			result.Error = errUserNotFound
		}
		results = append(results, result)
	}

	return results
}

func FromContext(ctx context.Context) *Loader {
	return ctx.Value(key).(*Loader)
}
