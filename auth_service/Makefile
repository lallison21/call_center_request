include scripts/*.mk

.PHONY: build run_docker create_migrate

docker_up:
	docker compose -f deployments/development/docker-compose.yml down -v
	docker compose -f deployments/development/docker-compose.yml up -d postgres
	docker compose -f deployments/development/docker-compose.yml up -d migration
	docker compose -f deployments/development/docker-compose.yml up -d --build --force-recreate auth_service
	docker compose -f deployments/development/docker-compose.yml ps

migrate_up:
	docker compose -f deployments/development/docker-compose.yml run migration up

docker_down:
	docker compose -f deployments/development/docker-compose.yml down -v

migrate_down:
	docker compose -f deployments/development/docker-compose.yml run migration down

create_migrate:
	migrate create -ext sql -dir migrations -seq ${MIGRATION_NAME}

.DEFAULT_GOAL: run_docker
