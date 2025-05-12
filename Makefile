BACKEND_NAME=lockbox-backend
FRONTEND_NAME=lockbox-frontend
POSTGRES_NAME=lockbox-postgres
MINIO_NAME=lockbox-minio

MIGRATIONS_DIR=backend/migrations
DB_URL=postgres://postgres:postgres@localhost:6969/lock_box?sslmode=disable
DB_CONTAINER_URL=postgres://postgres:postgres@${POSTGRES_NAME}:5432/lock_box?sslmode=disable

GREEN=\033[0;32m
NC=\033[0m

GREEN=\033[0;32m
NC=\033[0m

up:
	@echo "Запуск всех сервисов..."
	docker compose up -d

down:
	@echo "Остановка всех сервисов..."
	docker compose down

down_force:
	@echo "Полная очистка всех сервисов и образов..."
	docker compose down --volumes --remove-orphans
	docker-compose down --rmi all --volumes --remove-orphans

init: down_force build up
	@echo "Инициализация завершена"

init_deploy:
	@echo "Полный деплой с очисткой и пересборкой..."
	docker compose -f docker-compose.prod.yml down --volumes --remove-orphans
	docker compose -f docker-compose.prod.yml build --no-cache
	docker compose -f docker-compose.prod.yml up -d
	@echo "Инициализационный деплой завершён"

build:
	@echo "Сборка образов..."
	docker compose build --no-cache

restart: down up
	@echo "Сервисы перезапущены"

console_backend:
	docker exec -it ${BACKEND_NAME} sh

console_minio:
	docker exec -it ${MINIO_NAME} sh

frontend_logs:
	docker compose logs -f frontend

backend_logs:
	docker compose logs -f backend

db_logs:
	docker compose logs -f postgres

minio_logs:
	docker compose logs -f minio

migrate:
	@echo "Применение миграций..."
	goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" up

migrate_down:
	@echo "Откат миграций..."
	goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" down

migrate_status:
	@echo "Статус миграций..."
	goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" status

restart_db:
	@echo "База данных перезапущена"
	docker compose stop postgres
	docker compose rm -f postgres
	docker compose build --no-cache postgres
	docker compose up -d --no-deps postgres

help:
	@echo "Доступные команды:"
	@echo ""
	@echo "  make up             	- Запуск всех сервисов в фоне"
	@echo "  make down           	- Остановка всех сервисов"
	@echo "  make down_force     	- Полная очистка (контейнеры, volumes, образы)"
	@echo "  make init           	- Полная пересборка и запуск (down_force + build + up)"
	@echo "  make build          	- Пересобрать все образы без кеша"
	@echo "  make restart        	- Перезапуск сервисов (down + up)"
	@echo "  make restart_backend   - Перезапуск backend"
	@echo "  make restart_frontend  - Перезапуск сервисов (down + up)"
	@echo "  make test_frontend     - Запуск тестов фронтенда"
	@echo ""
	@echo "  make console        	- Войти в контейнер бэкенда"
	@echo "  make frontend_logs  	- Просмотр логов фронтенда"
	@echo "  make backend_logs   	- Просмотр логов бэкенда"
	@echo "  make db_logs        	- Просмотр логов PostgreSQL"
	@echo ""
	@echo "  make migrate        	- Применить миграции БД"
	@echo "  make migrate_down   	- Откатить последнюю миграцию БД"
	@echo "  make migrate_status 	- Показать статус миграций"
	@echo ""
	@echo "  make up_db          	- Запустить только PostgreSQL"
	@echo "  make wait_db        	- Ожидание готовности БД"
	@echo "  make restart_db       	- Перезапуск БД (down + up + wait)"
	@echo ""
	@echo "  make help           	- Показать это сообщение"
	@echo ""

restart_backend:
	@echo "Пересборка и перезапуск только backend..."
	docker compose stop backend
	docker compose rm -f backend
	docker compose build --no-cache backend
	docker compose up -d --no-deps backend
	make restart_db
	@echo "Backend пересобран и перезапущен"

restart_minio:
	@echo "Пересборка и перезапуск только minio..."
	docker compose stop minio
	docker compose rm -f minio
	docker compose build --no-cache minio
	docker compose up -d --no-deps minio
	@echo "minio пересобран и перезапущен"

restart_frontend:
	@echo "Пересборка и перезапуск только frontend..."
	docker compose stop frontend
	docker compose rm -f frontend
	docker compose build --no-cache frontend
	docker compose up -d --no-deps frontend
	@echo "Frontend пересобран и перезапущен"

test_frontend:
	@echo "Запуск тестов фронтенда..."
	docker exec $(FRONTEND_NAME) npm run types:check
	docker exec $(FRONTEND_NAME) npm run prettier:fix
	docker exec $(FRONTEND_NAME) npm run eslint:check
	docker exec $(FRONTEND_NAME) npm run build
	make restart_frontend

.PHONY: up down down_force init build restart console_backend frontend_logs backend_logs migrate migrate_down migrate_status restart_db help restart_frontend restart_backend test_frontend restart_minio minio_logs
