package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/ww24/graphql-tutorial/domain/service"
	"github.com/ww24/graphql-tutorial/presentation/graphql/dataloader"
	"github.com/ww24/graphql-tutorial/presentation/graphql/generated"
	"github.com/ww24/graphql-tutorial/presentation/graphql/resolver"
)

type server struct {
	resolver *resolver.Resolver
	user     service.User
	schedule service.Schedule
}

func newServer(
	resolver *resolver.Resolver,
	user service.User,
	schedule service.Schedule,
) *server {
	return &server{
		resolver: resolver,
		user:     user,
		schedule: schedule,
	}
}

func (s *server) newGraphQLServer() *handler.Server {
	srv := handler.New(generated.NewExecutableSchema(resolver.New(s.resolver)))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Println(err)
		debug.PrintStack()
		return errors.New("unexpected error")
	})
	return srv
}

func (s *server) dataloaderMiddleware(h http.Handler) http.Handler {
	return dataloader.Middleware(s.user)(h)
}
