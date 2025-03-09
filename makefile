PROJECT_NAME=home_server
SERVICE_NAME=app

up:
	docker-compose up -d

down:
	docker-compose down

down_force:
	docker-compose down --volumes --remove-orphans
	docker rmi $(PROJECT_NAME)

init:
	docker-compose down --volumes --remove-orphans
	docker-compose build
	docker-compose up -d

console:
	docker exec -it $$(docker-compose ps -q $(SERVICE_NAME)) sh

restart:
	docker-compose down --remove-orphans
	docker-compose build
	docker-compose up -d

.PHONY: up down down_force init console restart
