.PHONY: help build up down exec
.DEFAULT_GOAL := help

build:	## docker build
				docker compose build

up:	## docker compose up
				docker compose up -d

down: ## docker compose down
				docker compose down

exec: ## docker compose exec
				docker compose exec app bash

help:	## Show options
				@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
								awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'