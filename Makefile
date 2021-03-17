
PHONY: gqlgen
gqlgen:
	gqlgen generate

PHONY: run
run:
	go run ./cmd/schedule
