#!/usr/bin/env bash

BASE_DATASTORE_PATH="scripts/migrations/postgresql/go-migrate/urls"

ENVIRONMENT_NAME=${1?"A migration domain name should be provided!"}

declare -A DB_URLS=(
    ["playground"]="$BASE_DATASTORE_PATH/playground-src-url.sh"
    ["integration-test"]="$BASE_DATASTORE_PATH/integration-test-src-url.sh"
)

${DB_URLS[$ENVIRONMENT_NAME]}
