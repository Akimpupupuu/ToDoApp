include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	docker compose up -d postgres

env-down:
	docker compose down postgres

env-cleanup:
	@read -p "Clear all volume files of env? Warning!. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down postgres && \
		sudo rm -rf out/pgdata && \
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