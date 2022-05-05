#!/usr/bin/env bash

MIGRATOR_PATH="./scripts/migrations/postgresql/go-migrate/run/migrate.sh"

declare -a DOMAINS=(
    "auth"
    "manager-devices"
    "manager-bank-accounts"
)

DB_ENV_NAME=$1
if [[ -z $MIGRATION_PATH ]]; then
    echo "DB_ENV_NAME is empty - setting default to playground"
    DB_ENV_NAME="playground"
fi

for domain in "${DOMAINS[@]}"; do
    # call the migration script here
    echo "Domain migration => $domain"
    $MIGRATOR_PATH "$domain" $DB_ENV_NAME up
done
