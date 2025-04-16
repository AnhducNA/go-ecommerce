GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING = "root:password@(127.0.0.1:3306)/ecommerce?charset=utf8mb4"
GOOSE_MIGRATION_DIR = "./migrations"

APP_NAME = server

hello: 
	echo "Hello"

dev:
run: 
	go run ./cmd/${APP_NAME}/main.go

create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

upse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=$(GOOSE_MIGRATION_DIR) up
downse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=$(GOOSE_MIGRATION_DIR) down
resetse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlgen:
	sqlc generate

.PHONY: run downse upse resetse

.PHONY: air
