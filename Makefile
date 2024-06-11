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

dsn = "user=postgres dbname=postgres password=postgres host=localhost port=5432 sslmode=disable"

migration_up:
	goose -dir ./migrations postgres $(dsn) up

migration_status:
	goose -dir ./migrations postgres $(dsn) status

migration_down:
	goose -dir ./migrations postgres $(dsn) down
