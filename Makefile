include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d postgres

env-down:
	@docker compose down postgres

env-cleanup:
	@read -p "Clear all volume files of env? Warning!. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down postgres port-forwarder && \
		sudo rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Env files are cleared"; \
	else \
		echo "Cleaning is canceled"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Parametr seq is empty"; \
		exit 1; \
	fi;
	docker compose run --rm postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down
	
migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Parametr action is empty"; \
		exit 1; \
	fi;
	docker compose run --rm postgres-migrate \
		-path /migrations \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable \
		"$(action)"

env-port-forwarder:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

logs-cleanup:
	@read -p "Clear all log files? Warning!. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "Log files are cleared"; \
	else \
		echo "Cleaning is canceled"; \
	fi

todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todoapp/main.go