.PHONY:
	build-migrator migrate-create migrate-up migrate-down setup-postgres start-postgres stop-postgres

GOOSE_MIGRATOR_BINARY_NAME = goose-migrator
GOOSE_MIGRATOR_DIR = migrator

DATABASE_USER ?= root
DATABASE_PASSWORD ?= password
DATABASE_NAME ?= postgres
DATABASE_HOST ?= localhost
DATABASE_PORT ?= 5432

DB_ENV_VARS = DATABASE_USER=${DATABASE_USER} \
         DATABASE_PASSWORD=${DATABASE_PASSWORD} \
         DATABASE_NAME=${DATABASE_NAME} \
         DATABASE_HOST=${DATABASE_HOST} \
         DATABASE_PORT=${DATABASE_PORT}

build-migrator:
	@cd $(GOOSE_MIGRATOR_DIR) && CGO_ENABLED=0 go build -o $(GOOSE_MIGRATOR_BINARY_NAME)

migrate-create: build-migrator
	@cd $(GOOSE_MIGRATOR_DIR) && $(DB_ENV_VARS) ./$(GOOSE_MIGRATOR_BINARY_NAME) create $(name) go

migrate-up: build-migrator
	@cd $(GOOSE_MIGRATOR_DIR) && $(DB_ENV_VARS) ./$(GOOSE_MIGRATOR_BINARY_NAME) up 

migrate-down: build-migrator
	@cd $(GOOSE_MIGRATOR_DIR) && $(DB_ENV_VARS) ./$(GOOSE_MIGRATOR_BINARY_NAME) down

setup-postgres:
	@docker run --name go-goose-migrator-example-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine

start-postgres:
	@docker start go-goose-migrator-example-postgres

stop-postgres:
	@docker stop go-goose-migrator-example-postgres