.PHONY: up down build migrate rollback run test fmt

up:
#	docker compose up -d postgres

down:
#	docker compose down

build:
#	docker compose build app
	cd app && go build ./...

migrate:
	# docker compose --profile tools run --rm migrate db:migrate
	cd migrate && bundle exec rake db:migrate

rollback:
	#	docker compose --profile tools run --rm migrate db:rollback
	cd migrate && bundle exec rake db:rollback

run:
#	docker compose up -d --build app
	cd app && go run ./cmd/server

test:
	cd app && go test ./...

fmt:
	cd app && gofmt -w $$(find . -name '*.go' -type f)
