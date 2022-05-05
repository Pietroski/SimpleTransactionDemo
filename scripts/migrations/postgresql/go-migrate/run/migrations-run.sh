#!/usr/bin/env bash

MIGRATIONS_PATH=$1
DB_SOURCE_URL=$2
TYPE=$3
OFFSET=$4

# DB_SOURCE_URL=$(./def-src-url.run)
if [[ "$DB_SOURCE_URL" -eq '' ]]; then
    DB_SOURCE_URL=$(./scripts/migrations/postgresql/go-migrate/run/def-src-url.sh) # called from the Makefile
fi

MIGRATE=migrate
# shellcheck disable=SC2206
declare -a MIGRATE_BASE_COMMAND=(${MIGRATE} -path ${MIGRATIONS_PATH} -database ${DB_SOURCE_URL} -verbose)

echo "Migrations -=> $TYPE -- $OFFSET"

case "$TYPE" in
'init')
    echo "Versioning migrations++"
    ${MIGRATE} create -ext sql -dir "${MIGRATIONS_PATH}" -seq init_schema
    ;;
'up')
    if [[ "$OFFSET" -gt 0 ]]; then
        echo "Migrating up..."
        "${MIGRATE_BASE_COMMAND[@]}" up "${OFFSET}"
    else
        echo "Migrating up..."
        "${MIGRATE_BASE_COMMAND[@]}" up
    fi
    ;;
'down')
    if [[ "$OFFSET" -gt 0 ]]; then
        echo "Migrating up..."
        "${MIGRATE_BASE_COMMAND[@]}" down "${OFFSET}"
    else
        echo "Migrating up..."
        "${MIGRATE_BASE_COMMAND[@]}" down
    fi
    ;;
esac
