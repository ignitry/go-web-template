.PHONY: down build setup migrate rollback run run-production test fmt

down:
	overmind quit

build:
	cd app && go build -trimpath -o ../bin/server ./cmd/server

setup:
	bundle exec rake db:setup

migrate:
	bundle exec rake db:migrate

rollback:
	bundle exec rake db:rollback

run:
	@if overmind status >/dev/null 2>&1; then overmind restart web; else overmind start -D -f Procfile; fi

run-production:
	overmind start -f Procfile.production --auto-restart all

test:
	cd app && go test ./...

fmt:
	cd app && gofmt -w $$(find . -name '*.go' -type f)
