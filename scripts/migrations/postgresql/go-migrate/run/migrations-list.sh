#!/usr/bin/env bash

BASE_DATASTORE_PATH="internal/adaptors/datastore/postgresql"
MIGRATIONS="migrations"

MANAGER="manager"

MIGRATION_NAME=${1?"A migration domain name should be provided!"}

declare -A migrations=(
       [auth]="$BASE_DATASTORE_PATH/auth/$MIGRATIONS"

       [manager-devices]="$BASE_DATASTORE_PATH/$MANAGER/devices/$MIGRATIONS"
       [manager-bank-accounts]="$BASE_DATASTORE_PATH/$MANAGER/bank-accounts/$MIGRATIONS"

       [playground-backfil]="playground/services/postgresql"
)

echo "${migrations[$MIGRATION_NAME]}"
