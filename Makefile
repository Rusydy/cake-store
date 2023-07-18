# load env
include .env

.PHONY: run
run:
	@go run cmd/main.go

.PHONY: dc-run
dc-run:
	@docker-compose up --build

.PHONY: migrate-create
migrate-create:
ifndef name
	$(error name is not set. Please specify the migration name using 'make migrate-create name=<migration_name>')
endif
	@migrate create -ext sql -dir internal/database/migrations -seq $(name)

.PHONY: migrate-up
migrate-up:
	@migrate -path internal/database/migrations -database $(DATABASE_URL) up

.PHONY: lint
lint:
	@golangci-lint run ./...