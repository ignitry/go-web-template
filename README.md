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

## Migration naming

Always prefix migration filenames with a 14-digit timestamp in `YYYYMMDDhhmmss` format (24-hour time), followed by an underscore and a descriptive name. Do not use sequential numbers such as `001`.

Example: `20261115100630_create_example_records.rb`.

Use the creation timestamp for each new migration and keep prefixes unique so migrations run in chronological order.

## Rename project

```sh
scripts/rename.sh example.com/my-project my_project_development
```

The first argument sets the Go module and project name. The second sets the PostgreSQL database name in Docker Compose and migration defaults.
