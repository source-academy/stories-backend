#!/bin/bash

set -euo pipefail

export DB_HOSTNAME="$PG_HOST"
export DB_USERNAME="$PG_USER"
export DB_NAME="$PG_DATABASE"

while ! pg_isready -q -h "$PG_HOST" -U "$PG_USER"; do
    echo "$(date) - waiting for database to start"
    sleep 0.5
done

# Create database if it doesn't exist.
if [[ -z `psql -h "$PG_HOST" -U "$PG_USER" -lt | cut -d '|' -f 1 | grep -w "$PG_DATABASE"` ]]; then
    echo "Database $PG_DATABASE does not exist. Creating..."
    make db_create
fi

make db_migrate
