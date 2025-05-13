BACKEND_NAME=lockbox-backend
FRONTEND_NAME=lockbox-frontend
POSTGRES_NAME=lockbox-postgres
MINIO_NAME=lockbox-minio

MIGRATIONS_DIR=backend/migrations
DB_URL=postgres://postgres:postgres@localhost:6969/lock_box?sslmode=disable
DB_CONTAINER_URL=postgres://postgres:postgres@${POSTGRES_NAME}:5432/lock_box?sslmode=disable

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
	@printf "\nДоступные команды:\n\n"
	@printf "  ${GREEN}make up               			${NC}- Запуск всех сервисов в фоне\n"
	@printf "  ${GREEN}make down             			${NC}- Остановка всех сервисов\n"
	@printf "  ${GREEN}make down_force       			${NC}- Полная очистка (контейнеры, volumes, образы)\n"
	@printf "  ${GREEN}make init             			${NC}- Полная пересборка и запуск (down_force + build + up)\n"
	@printf "  ${GREEN}make build            			${NC}- Пересобрать все образы без кеша\n"
	@printf "  ${GREEN}make restart          			${NC}- Перезапуск сервисов (down + up)\n"
	@printf "  ${GREEN}make restart_backend  			${NC}- Перезапуск backend\n"
	@printf "  ${GREEN}make restart_frontend 			${NC}- Перезапуск frontend\n"
	@printf "  ${GREEN}make test_frontend    			${NC}- Запуск тестов фронтенда\n"
	@printf "\n"
	@printf "  ${GREEN}make console_backend				${NC}- Войти в контейнер бэкенда\n"
	@printf "  ${GREEN}make console_minio				${NC}- Войти в контейнер minio\n"
	@printf "  ${GREEN}make frontend_logs    			${NC}- Просмотр логов фронтенда\n"
	@printf "  ${GREEN}make backend_logs     			${NC}- Просмотр логов бэкенда\n"
	@printf "  ${GREEN}make db_logs          			${NC}- Просмотр логов PostgreSQL\n"
	@printf "\n"
	@printf "  ${GREEN}make migrate          			${NC}- Применить миграции БД\n"
	@printf "  ${GREEN}make migrate_down     			${NC}- Откатить последнюю миграцию БД\n"
	@printf "  ${GREEN}make migrate_status   			${NC}- Показать статус миграций\n"
	@printf "\n"
	@printf "  ${GREEN}make up_db            			${NC}- Запустить только PostgreSQL\n"
	@printf "  ${GREEN}make wait_db          			${NC}- Ожидание готовности БД\n"
	@printf "  ${GREEN}make restart_db       			${NC}- Перезапуск БД (down + up + wait)\n"
	@printf "  ${GREEN}make restart_minio				${NC}- Перезапуск minio \n"
	@printf "\n"
	@printf "  ${GREEN}make help             			${NC}- Показать это сообщение\n\n"

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
