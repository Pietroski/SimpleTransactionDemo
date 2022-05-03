#!/usr/bin/env bash

MIGRATIONS_PATH=$1
DB_DATA_SOURCE_NAME=$2
TYPE=$3
OFFSET=$4

MIGRATE=migrate

echo "Migrations -=> $TYPE -- $OFFSET"

case "$TYPE" in
    'init')
        echo "Versioning migrations++"
        ${MIGRATE} create -ext sql -dir "${MIGRATIONS_PATH}" -seq init_schema
    ;;
    'up')
        if [[ "$OFFSET" -gt 0 ]]; then
            echo "Migrating up..."
            ${MIGRATE} -path "${MIGRATIONS_PATH}" -database "${DB_DATA_SOURCE_NAME}" -verbose up "${OFFSET}"
        else
            echo "Migrating up..."
            ${MIGRATE} -path "${MIGRATIONS_PATH}" -database "${DB_DATA_SOURCE_NAME}" -verbose up
        fi
    ;;
    'down')
        if [[ "$OFFSET" -gt 0 ]]; then
            echo "Migrating up..."
            ${MIGRATE} -path "${MIGRATIONS_PATH}" -database "${DB_DATA_SOURCE_NAME}" -verbose down "${OFFSET}"
        else
            echo "Migrating up..."
            ${MIGRATE} -path "${MIGRATIONS_PATH}" -database "${DB_DATA_SOURCE_NAME}" -verbose down
        fi
esac
