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

# Keep the display name readable when the project name contains separators.
display_name=$(printf '%s\n' "$project" | awk '{ gsub("[-_]", " "); $1=toupper(substr($1,1,1)) substr($1,2); print }')

# Rename every text file so Go packages, tests, templates, and documentation
# cannot retain references that are missed by an explicit file list.
find "$root" -type f ! -path "$root/.git/*" ! -path "$root/scripts/rename.sh" \
  -exec sh -c 'grep -Iq . "$1" && exit 0 || exit 1' sh {} \; \
  -exec sh -c '
    file=$1
    temporary="$file.tmp"
    MODULE_NAME="$2" PROJECT_NAME="$3" DISPLAY_NAME="$4" DATABASE_NAME="$5" awk '\''{
      gsub("starter_development", ENVIRON["DATABASE_NAME"])
      gsub("Starter", ENVIRON["DISPLAY_NAME"])
      gsub(/"starter\//, "\"" ENVIRON["MODULE_NAME"] "/")
      gsub("starter", ENVIRON["PROJECT_NAME"])
      print
    }'\'' "$file" > "$temporary"
    mv "$temporary" "$file"
  ' sh {} "$module" "$project" "$display_name" "$database" \;

temporary="$root/app/go.mod.tmp"
awk -v module="$module" '$1 == "module" { print "module " module; next } { print }' "$root/app/go.mod" > "$temporary"
mv "$temporary" "$root/app/go.mod"
