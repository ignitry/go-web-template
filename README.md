# Go Web Template

A template for web development in Go with PostgreSQL database.

## Requirements

- Docker Compose

## Start

```sh
cp .env.example .env
make migrate
make run
```

Open http://localhost:8080/healthz.

## Commands

```sh
make up
make down
make build
make migrate
make rollback
make run
make fmt
```

Migrations are managed only with ActiveRecord under `migrate/db/migrate`.
Docker Compose reads `.env` automatically.

## Rename project

```sh
scripts/rename.sh example.com/my-project my_project_development
```

The first argument sets the Go module and project name. The second sets the PostgreSQL database name in Docker Compose and migration defaults.
