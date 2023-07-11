#!/bin/bash

# Usage: ./scripts/check-migrations-ordering.sh <base-ref>
# This script checks that all migrations in the current branch are
# newer than the last merged migration in the base branch.
# Arguments:
#   base-ref: Required. The branch/ref to compare against.
# Example:
#   ./scripts/check-migrations-ordering.sh main

set -euo pipefail

BASE_REF=$1

# Gets the list of migrations in the specified branch, without the folder name prefix
get_migrations_in_branch() {
    git ls-tree -r --name-only $1 migrations | sed 's/^migrations\///' | sort
}

# Get the list of migrations in base branch
echo "Checking migrations in ref $BASE_REF..."
BASE_MIGRATIONS=$(get_migrations_in_branch $BASE_REF)

# Get the list of migrations in current branch
CURRENT_MIGRATIONS=$(get_migrations_in_branch HEAD)

# Get the list of migrations that are in the current branch but not in the base branch
NEW_MIGRATIONS=$(comm -23 <(echo "$CURRENT_MIGRATIONS") <(echo "$BASE_MIGRATIONS"))
if [ -z "$NEW_MIGRATIONS" ]; then
    echo "No new migrations found in current branch, exiting..."
    exit 0
fi
echo "Found the following NEW migrations:"
# Indent the list of migrations
echo "$(echo "$NEW_MIGRATIONS" | sed 's/^/    /')"

echo "Checking timestamps..."
# Get the timestamp of the last migration in the base branch
LAST_MIGRATION_TIMESTAMP=$(echo "$BASE_MIGRATIONS" | tail -n 1 | sed 's/[^0-9]*\([0-9]*\).*/\1/')
echo "Last known migration timestamp: ${LAST_MIGRATION_TIMESTAMP:-"none"}"
if [ -z "$LAST_MIGRATION_TIMESTAMP" ]; then
    echo "No migrations found in base branch, all migrations allowed..."
    exit 0
fi

TIMESTAMPS_TO_CHECK=$(echo "$NEW_MIGRATIONS" | sed 's/[^0-9]*\([0-9]*\).*/\1/')
for TIMESTAMP in $TIMESTAMPS_TO_CHECK; do
    if [ "$TIMESTAMP" -lt "$LAST_MIGRATION_TIMESTAMP" ]; then
        echo "ERROR: Migration with timestamp $TIMESTAMP is older than the last merged migration timestamp $LAST_MIGRATION_TIMESTAMP"
        exit 1
    fi
done
echo "All migrations are newer than the last merged migration timestamp, exiting..."
