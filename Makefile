build:
	docker-compose build

build_up:
	docker-compose up --build

up:
	docker-compose up

down:
	docker-compose down

exec:
	docker-compose exec backend sh

exec_db:
	docker-compose exec db bash
