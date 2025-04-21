BACKEND_NAME=lockbox-backend
FRONTEND_NAME=lockbox-frontend
POSTGRES_NAME=lockbox-postgres

MIGRATIONS_DIR=backend/migrations
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

init: down_force build up
	@echo "Инициализация завершена"

init_deploy:
	@echo "Полный деплой с очисткой и пересборкой..."
	docker compose -f docker-compose.prod.yml down --volumes --remove-orphans
	docker compose -f docker-compose.prod.yml build --no-cache
	docker compose -f docker-compose.prod.yml up -d
	@echo "✅ Инициализационный деплой завершён"

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
help:
	@echo "Доступные команды:"
	@echo ""
	@echo "  make up             - Запуск всех сервисов в фоне"
	@echo "  make down           - Остановка всех сервисов"
	@echo "  make down_force     - Полная очистка (контейнеры, volumes, образы)"
	@echo "  make init           - Полная пересборка и запуск (down_force + build + up)"
	@echo "  make build          - Пересобрать все образы без кеша"
	@echo "  make restart        - Перезапуск сервисов (down + up)"
	@echo ""
	@echo "  make console        - Войти в контейнер бэкенда"
	@echo "  make frontend_logs  - Просмотр логов фронтенда"
	@echo "  make backend_logs   - Просмотр логов бэкенда"
	@echo "  make db_logs        - Просмотр логов PostgreSQL"
	@echo ""
	@echo "  make migrate        - Применить миграции БД"
	@echo "  make migrate_down   - Откатить последнюю миграцию БД"
	@echo "  make migrate_status - Показать статус миграций"
	@echo ""
	@echo "  make up_db          - Запустить только PostgreSQL"
	@echo "  make wait_db        - Ожидание готовности БД"
	@echo "  make reset_db       - Перезапуск БД (down + up + wait)"
	@echo ""
	@echo "  make help           - Показать это сообщение"
	@echo ""

.PHONY: up down down_force init build restart console frontend_logs backend_logs db_logs migrate migrate_down migrate_status up_db wait_db reset_db help
