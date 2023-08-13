#!/bin/bash

set -euo pipefail

SCRIPT_DIR="$(dirname "$(realpath "${BASH_SOURCE[0]}")")"

COMPOSE_FILE_PATH="$SCRIPT_DIR/docker-compose.yml"
echo "Using Docker Compose file: $COMPOSE_FILE_PATH"

BACKUP_DIR="./backups"
mkdir -p "$BACKUP_DIR"

LAST_BACKUP_FILENAME="$(ls -1 "$BACKUP_DIR" | sort | tail -1)"
LAST_BACKUP_FILE_PATH="$BACKUP_DIR/$LAST_BACKUP_FILENAME"

NEW_BACKUP_FILENAME="pgdump_`date +%Y%m%d-%H%M%S`.sql"
BACKUP_FILE_PATH="$BACKUP_DIR/$NEW_BACKUP_FILENAME"

BACKUP_COMMAND="'pg_dumpall -c -h \$PG_HOST -U \$PG_USER'"
SQL_DUMP="$(docker compose -f "$COMPOSE_FILE_PATH" run --rm -it --entrypoint "bash -c $BACKUP_COMMAND" migrator)"

if [[ -z "$LAST_BACKUP_FILENAME" ]]; then
    echo "No previous backups found, creating new backup..."
else
    echo "Last backup found: $LAST_BACKUP_FILE_PATH"
    if [[ -z `echo "$SQL_DUMP" | diff "$LAST_BACKUP_FILE_PATH" -` ]]; then
        echo "No changes since last backup, exiting..."
        exit 0
    else
        echo "Changes found since last backup, saving new backup..."
    fi
fi

echo "Backup file path: $BACKUP_FILE_PATH"
echo "$SQL_DUMP" > "$BACKUP_FILE_PATH"
echo "Backup saved successfully!"
