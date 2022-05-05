#!/usr/bin/env bash

BASE_DATASTORE_PATH="internal/adaptors/datastore/postgresql"
MIGRATIONS="migrations"

MANAGER="manager"

MIGRATOR_PATH="./scripts/migrations/postgresql/go-migrate/run/migrations-run.sh"

declare -A migrations=(
       [auth]="$BASE_DATASTORE_PATH/auth/$MIGRATIONS"

       [manager-devices]="$BASE_DATASTORE_PATH/$MANAGER/devices/$MIGRATIONS"
       [manager-bank-accounts]="$BASE_DATASTORE_PATH/$MANAGER/bank-accounts/$MIGRATIONS"
)

for ctx in "${!migrations[@]}"; do
    # call the migration script here
    echo "$ctx => ${migrations[$ctx]}"
    $MIGRATOR_PATH "${migrations[$ctx]}" "" up
done
