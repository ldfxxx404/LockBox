BACKEND_NAME=lockbox-backend
FRONT_NAME=lockbox-frontend
POSTGRES_NAME=postgres

MIGRATIONS_DIR=back/migrations
DB_URL=postgres://postgres:postgres@localhost:6969/lock_box?sslmode=disable

up:
	docker compose up -d

down:
	docker compose down

down_force:
	docker compose down --volumes --remove-orphans
	docker rmi -f $$(docker images -q $(BACKEND_NAME))
	docker rmi -f $$(docker images -q $(POSTGRES_NAME))
	docker rmi -f $$(docker images -q $(FRONT_NAME))

init:
	docker compose down --volumes --remove-orphans
	docker compose build
	docker compose up -d

console:
	docker exec -it $$(docker compose ps -q $(BACKEND_NAME)) sh

restart:
	docker compose down --remove-orphans
	docker compose build
	docker compose up -d

migrate:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

migrate_down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

up_db:
	docker-compose up -d --no-deps $(POSTGRES_NAME)

.PHONY: up down down_force init console restart migrate migrate_down up_db
