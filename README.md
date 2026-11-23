# Go Web Template

A template for web development in Go with PostgreSQL database.

## Requirements

- Ubuntu or Debian
- Go 1.22.2 or newer
- Ruby 3.2 or newer and Bundler (ActiveRecord 8 requires Ruby 3.2+)
- PostgreSQL, libpq-dev, build-essential, and ruby-dev matching your Ruby version
- Overmind and tmux

Install Go, Ruby, Bundler, and Overmind before continuing. Use a supported Ruby
version rather than an older distribution default.

```sh
sudo apt-get update
sudo apt-get install -y postgresql libpq-dev build-essential tmux
sudo systemctl enable --now postgresql
sudo -u postgres createuser --login --createdb --pwprompt starter
bundle install
cp .env.example .env
make setup
make run
```

Set the PostgreSQL role password to match DATABASE_URL in .env.
Open http://localhost:8080/healthz.

## Commands

```sh
make build
make setup
make migrate
make rollback
make run
make down
make test
make fmt
```

Run commands from the repository root. Migrations use the root Gemfile and
Rakefile, with ActiveRecord migrations under `migrate/db/migrate`. `make setup`
creates the configured database and applies migrations. Database creation
requires a PostgreSQL role with CREATEDB permission.

Overmind loads the root .env and starts the web process using Procfile.
`make run` starts it in the background or restarts an existing web process.
`make down` stops the application. PostgreSQL runs as a host service.
Rake loads the same .env through dotenv. Existing process environment variables
take precedence over .env values. Local database connections use localhost.

Run the Go test suite with `make test`.

## Production

Install the same host dependencies, run `bundle install`, and configure .env
with APP_ENV=production and the production DATABASE_URL and HTTP_PORT.
Provision the database and role before migrating.

```sh
make build
make migrate
make run-production
```

The binary is built at `bin/server`. Procfile.production runs it through
Overmind with automatic process restart. Run `make run-production` under a
host service manager for startup on boot and supervisor recovery. Keep .env
on the host and out of version control.

## Migration naming

Always prefix migration filenames with a 14-digit timestamp in `YYYYMMDDhhmmss` format (24-hour time), followed by an underscore and a descriptive name. Do not use sequential numbers such as `001`.

Example: `20261115100630_create_example_records.rb`.

Use the creation timestamp for each new migration and keep prefixes unique so migrations run in chronological order.

## Rename project

```sh
scripts/rename.sh example.com/my-project my_project_development
```

The first argument sets the Go module and project name. The second sets the
PostgreSQL database name in environment examples and migration defaults.
