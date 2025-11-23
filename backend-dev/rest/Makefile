ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

.PHONY: dev
dev:
	@air

.PHONY: test
test:
	@go test -v ./...

MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: migration
migration:
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@go run cmd/migrate/main.go up

.PHONY: migrate-down
migrate-down:
	@go run cmd/migrate/main.go down

LOCAL_POSTGRES_DOCKER_PATH = tools/docker-compose.yml

.PHONY: postgres
postgres:
	docker-compose -f $(LOCAL_POSTGRES_DOCKER_PATH) up -d
