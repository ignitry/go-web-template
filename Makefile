.PHONY: up down build migrate rollback run fmt

up:
	docker compose up -d postgres

down:
	docker compose down

build:
	docker compose build app

migrate:
	docker compose --profile tools run --rm migrate db:migrate

rollback:
	docker compose --profile tools run --rm migrate db:rollback

run:
	docker compose up -d --build app

fmt:
	cd app && gofmt -w $$(find . -name '*.go' -type f)
