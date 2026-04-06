include .env
export

export ROOT=$(shell pwd)

db-up:
	@docker compose up -d db

db-down:
	@docker compose down db

migrate-up:
	@docker compose up --build migrator

migrate-down:
	@docker compose run --rm migrator ./migrator-app down

run-petstore: 
	@go run ./cmd/main.go

lint: 
	@golangci-lint run

lint-fix: 
	@golangci-lint run --fix