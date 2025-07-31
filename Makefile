# --- env --- 
SHELL := /bin/bash

DOCKER_COMPOSE := @docker compose

BACKEND_NAME=lockbox-backend
FRONTEND_NAME=lockbox-frontend
POSTGRES_NAME=lockbox-postgres
MINIO_NAME=lockbox-minio

MIGRATIONS_DIR=migrations
DB_URL=postgres://postgres:postgres@postgres:5432/lock_box?sslmode=disable
DB_CONTAINER_URL=postgres://postgres:postgres@${POSTGRES_NAME}:5432/lock_box?sslmode=disable

# --- misc ---
up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

down_force:
	$(DOCKER_COMPOSE) down --volumes --remove-orphans
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans

init: down_force build up
	@echo "Инициализация завершена"
	@echo "backend docs http://localhost:5000/docs/index.html"
	@echo "frontend http://localhost:3000"

init_prod: down_force 
	$(DOCKER_COMPOSE) -f docker-compose.prod.yml build --no-cache
	$(DOCKER_COMPOSE) -f docker-compose.prod.yml up -d

build:
	$(DOCKER_COMPOSE) build --no-cache

restart: restart_backend restart_db restart_frontend restart_minio

# --- console ---
console_bd:
	@docker exec -it ${POSTGRES_NAME} ${SHELL}

console_backend:
	@docker exec -it ${BACKEND_NAME} ${SHELL}

console_minio:
	@docker exec -it ${MINIO_NAME} ${SHELL}

console_front:
	@docker exec -it ${FRONTEND_NAME} ${SHELL}

# --- logs ---
frontend_logs:
	$(DOCKER_COMPOSE) logs -f frontend

backend_logs:
	$(DOCKER_COMPOSE) logs -f backend

db_logs:
	$(DOCKER_COMPOSE) logs -f postgres

minio_logs:
	$(DOCKER_COMPOSE) logs -f minio

# --- db, migrate ---
migrate:
	@docker exec ${BACKEND_NAME} goose -dir migrations postgres ${DB_URL} up

migrate_down:
	@docker exec ${BACKEND_NAME} goose -dir migrations postgres ${DB_URL} down

migrate_status:
	@docker exec ${BACKEND_NAME} goose -dir migrations postgres ${DB_URL} status

dump:
	./backend/scripts/dump.sh 

restore: 
	./backend/scripts/restore_dump.sh 

# --- docs ---
swag_generate:
	@cd backend && \
	./scripts/swag.sh

help:
	@echo "make up                              - Запуск всех сервисов в фоне"
	@echo "make down                            - Остановка всех сервисов"
	@echo "make down_force                      - Полная очистка (контейнеры, volumes, образы)"
	@echo "make init                            - Полная пересборка и запуск (down_force + build + up)"
	@echo "make build                           - Пересобрать все образы без кеша"
	@echo "make restart                         - Перезапуск сервисов (down + up)"
	@echo "make restart_backend                 - Перезапуск backend"
	@echo "make restart_frontend                - Перезапуск сервисов (down + up)"
	@echo "make test_frontend                   - Запуск тестов фронтенда"
	@echo "make test_backend                    - Запуск тестов бэкенда"
	@echo "make console                         - Войти в контейнер бэкенда"
	@echo "make frontend_logs                   - Просмотр логов фронтенда"
	@echo "make backend_logs                    - Просмотр логов бэкенда"
	@echo "make db_logs                         - Просмотр логов PostgreSQL"
	@echo "make migrate                         - Применить миграции БД"
	@echo "make migrate_down                    - Откатить последнюю миграцию БД"
	@echo "make migrate_status                  - Показать статус миграций"
	@echo "make up_db                           - Запустить только PostgreSQL"
	@echo "make wait_db                         - Ожидание готовности БД"
	@echo "make restart_db                      - Перезапуск БД (down + up + wait)"
	@echo "make help                            - Показать это сообщение"

# --- restarts ---
restart_backend:
	$(DOCKER_COMPOSE) stop backend
	$(DOCKER_COMPOSE) rm -f backend
	$(DOCKER_COMPOSE) build --no-cache backend
	$(DOCKER_COMPOSE) up -d --no-deps backend
	@make restart_db

restart_minio:
	$(DOCKER_COMPOSE) stop minio
	$(DOCKER_COMPOSE) rm -f minio
	$(DOCKER_COMPOSE) build --no-cache minio
	$(DOCKER_COMPOSE) up -d --no-deps minio

restart_frontend:
	$(DOCKER_COMPOSE) stop frontend
	$(DOCKER_COMPOSE) rm -f frontend
	$(DOCKER_COMPOSE) build --no-cache frontend
	$(DOCKER_COMPOSE) up -d --no-deps frontend

restart_db:
	$(DOCKER_COMPOSE) stop postgres
	$(DOCKER_COMPOSE) rm -f postgres
	$(DOCKER_COMPOSE) build --no-cache postgres
	$(DOCKER_COMPOSE) up -d --no-deps postgres

# --- tests ---
test_frontend:
	@docker exec $(FRONTEND_NAME) yarn run types:check
	@docker exec $(FRONTEND_NAME) yarn run prettier:fix
	@docker exec $(FRONTEND_NAME) yarn run eslint:check
	@docker exec $(FRONTEND_NAME) yarn run build
	@make restart_frontend

test_backend:
	@cd backend && \
	./scripts/check.sh


.PHONY: up down down_force init build restart console_backend frontend_logs backend_logs migrate migrate_down migrate_status restart_db help restart_frontend restart_backend test_frontend restart_minio minio_logs console_front test_backend  
