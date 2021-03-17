package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
)

const (
	envProduction  = "prd"
	envDevelopment = "dev"
)

var (
	listenPort  = 8081
	environment = envDevelopment
)

func init() {
	port := os.Getenv("PORT")
	if p, err := strconv.Atoi(port); err == nil {
		listenPort = p
	}

	env := os.Getenv("ENV")
	switch env {
	case envProduction, envDevelopment:
		environment = env
	}
}

func main() {
	s := register()
	srv := s.newGraphQLServer()

	mux := http.NewServeMux()
	if environment == envDevelopment {
		srv.Use(extension.Introspection{})
		srv.Use(apollotracing.Tracer{})
		mux.Handle("/", playground.Handler("Schedule", "/query"))
	}
	mux.Handle("/query", s.dataloaderMiddleware(srv))

	httpSrv := &http.Server{
		Addr:    ":" + strconv.Itoa(listenPort),
		Handler: mux,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Printf("http server is closed: %v\n", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	log.Printf("signal is received: %s\n", <-sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Printf("error occurred while shutting down process: %v\n", err)
	}
}
