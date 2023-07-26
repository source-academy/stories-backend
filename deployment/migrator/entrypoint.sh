#!/bin/bash

set -euo pipefail

# Environment variables for Golang, must match .env.example
export DB_HOSTNAME="$PG_HOST"
export DB_USERNAME="$PG_USER"
export DB_NAME="$PG_DATABASE"

# Limit CPU and memory usage for the migrator
# to prevent crashing the deployment server.
export GOMAXPROCS="1"
export GOMEMLIMIT="256MiB"

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
