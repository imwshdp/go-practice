ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

.PHONY: dev
dev:
	@go run cmd/*.go

LOCAL_POSTGRES_DOCKER_PATH = tools/postgresql/docker-compose.yml

.PHONY: postgres-up
postgres-up:
	docker-compose -f $(LOCAL_POSTGRES_DOCKER_PATH) up -d

.PHONY: postgres-down
postgres-down:
	docker-compose -f $(LOCAL_POSTGRES_DOCKER_PATH)  down -v

.PHONY: migrate-up
migrate-up:
	goose up

.PHONY: sqlc-gen
sqlc-gen:
	sqlc generate
