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