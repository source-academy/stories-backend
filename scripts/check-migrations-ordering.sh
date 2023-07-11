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
get_timestamps() {
    # Migration file format: yyyymmddhhmmss-description.sql
    # We use `|| true;` to avoid the script failing if the grep returns no results
    awk -F'-' '{print $1}' $1 | { grep '^[0-9]\{14\}$' || true; }
}
# Get the timestamp of the last migration in the base branch
LAST_MIGRATION_TIMESTAMP=$(get_timestamps <(echo "$BASE_MIGRATIONS") | tail -n 1)
echo "Last merged migration timestamp: ${LAST_MIGRATION_TIMESTAMP:-"none"}"
if [ -z "$LAST_MIGRATION_TIMESTAMP" ]; then
    echo "No migrations found in base branch, all migrations allowed..."
    exit 0
fi

TIMESTAMPS_TO_CHECK=$(get_timestamps <(echo "$NEW_MIGRATIONS"))
for TIMESTAMP in $TIMESTAMPS_TO_CHECK; do
    if [ "$TIMESTAMP" -lt "$LAST_MIGRATION_TIMESTAMP" ]; then
        echo "ERROR: Migration with timestamp $TIMESTAMP is older than the last merged migration timestamp $LAST_MIGRATION_TIMESTAMP"
        exit 1
    fi
done
echo "All migrations are newer than the last merged migration timestamp, exiting..."
