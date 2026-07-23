#!/bin/sh
set -eu

if [ "$#" -ne 2 ] || [ -z "$1" ] || [ -z "$2" ]; then
  echo "usage: rename.sh <module-name> <database-name>" >&2
  exit 1
fi

module=$1
database=$2
case "$module" in
  *[!A-Za-z0-9._/-]* | /* | */ | *'//'*)
    echo "invalid module name: $module" >&2
    exit 1
    ;;
esac

case "$database" in
  [A-Za-z_]* ) ;;
  * )
    echo "invalid database name: $database" >&2
    exit 1
    ;;
esac

case "$database" in
  *[!A-Za-z0-9_]* )
    echo "invalid database name: $database" >&2
    exit 1
    ;;
esac

project=${module##*/}
case "$project" in
  *[!A-Za-z0-9_-]* | '')
    echo "invalid project name: $project" >&2
    exit 1
    ;;
esac

root=$(CDPATH= cd "$(dirname "$0")/.." && pwd)

replace() {
  file=$1
  temporary="$file.tmp"
  PROJECT_NAME=$project DATABASE_NAME=$database awk '{ gsub("starter_development", ENVIRON["DATABASE_NAME"]); gsub("starter", ENVIRON["PROJECT_NAME"]); print }' "$file" > "$temporary"
  mv "$temporary" "$file"
}

replace "$root/docker-compose.yml"
replace "$root/.env.example"
replace "$root/README.md"
replace "$root/app/cmd/server/main.go"
replace "$root/migrate/config/database.yml"

temporary="$root/app/go.mod.tmp"
awk -v module="$module" '$1 == "module" { print "module " module; next } { print }' "$root/app/go.mod" > "$temporary"
mv "$temporary" "$root/app/go.mod"
