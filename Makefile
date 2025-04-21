BACKEND_NAME=lockbox-backend
FRONTEND_NAME=lockbox-frontend
POSTGRES_NAME=lockbox_postgres

MIGRATIONS_DIR=back/migrations
DB_URL=postgres://postgres:postgres@localhost:6969/lock_box?sslmode=disable
DB_CONTAINER_URL=postgres://postgres:postgres@${POSTGRES_NAME}:5432/lock_box?sslmode=disable

up:
	@echo "Запуск всех сервисов..."
	docker compose up -d

down:
	@echo "Остановка всех сервисов..."
	docker compose down

down_force:
	@echo "Полная очистка всех сервисов и образов..."
	docker compose down --volumes --remove-orphans
	docker rmi -f $$(docker images -q ${BACKEND_NAME}) || true
	docker rmi -f $$(docker images -q ${POSTGRES_NAME}) || true
	docker rmi -f $$(docker images -q ${FRONTEND_NAME}) || true

init: down_force build up migrate
	@echo "Инициализация завершена"

build:
	@echo "Сборка образов..."
	docker compose build --no-cache

restart: down up
	@echo "Сервисы перезапущены"

console:
	docker exec -it $$(docker compose ps -q ${BACKEND_NAME}) sh

frontend_logs:
	docker compose logs -f frontend

backend_logs:
	docker compose logs -f backend

db_logs:
	docker compose logs -f postgres

migrate:
	@echo "Применение миграций..."
	goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" up

migrate_down:
	@echo "Откат миграций..."
	goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" down

migrate_status:
	@echo "Статус миграций..."
	goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" status

up_db:
	@echo "Запуск только PostgreSQL..."
	docker compose up -d --no-deps ${POSTGRES_NAME}

wait_db:
	@echo "Ожидание готовности PostgreSQL..."
	until docker exec $$(docker compose ps -q ${POSTGRES_NAME}) pg_isready; do sleep 1; done

reset_db: down_db up_db wait_db
	@echo "База данных перезапущена"

.PHONY: up down down_force init build restart console frontend_logs backend_logs db_logs migrate migrate_down migrate_status up_db wait_db reset_db
