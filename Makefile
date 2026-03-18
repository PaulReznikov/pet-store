include .env
export

export ROOT=$(shell pwd)

db-up:
	@docker compose up -d db

db-down:
	@docker compose down db

migrations-up:
	@docker compose up --build migrator

migrate-down:
	@docker compose run --rm migrator ./migrator-app down

pgadmin-up:
	@docker compose up -d pgadmin

pgadmin-down:
	@docker compose down pgadmin