#!/usr/bin/env bash

DESTINATION="internal/adaptors/datastore/postgresql/auth/migrations"
SCHEMA_NAME="create_accounts_models"

migrate create -ext sql -dir $DESTINATION -seq $SCHEMA_NAME
