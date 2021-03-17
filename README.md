# graphql-tutorial

## Requirements

- Go 1.16.x


## Start server

```
make run
```

## Directory structure

```
.
├── Makefile
├── README.md
├── cmd
│   └── schedule
│       ├── main.go
│       ├── server.go
│       ├── wire.go
│       └── wire_gen.go
├── domain
│   ├── repository
│   │   ├── schedule.go
│   │   └── user.go
│   └── service
│       ├── schedule.go
│       ├── service.go
│       └── user.go
├── go.mod
├── go.sum
├── gqlgen.yml
├── infra
│   └── db
│       ├── db.go
│       ├── schedule.go
│       └── user.go
└── presentation
    └── graphql
        ├── dataloader
        │   └── dataloader.go
        ├── generated
        │   ├── federation.go
        │   └── generated.go
        ├── model
        │   ├── models_gen.go
        │   └── user.go
        ├── resolver
        │   ├── mutation.go
        │   ├── query.go
        │   ├── resolver.go
        │   ├── schedule.go
        │   └── user.go
        └── schema
            ├── mutation.graphqls
            ├── query.graphqls
            ├── schedule.graphqls
            └── user.graphqls
```
