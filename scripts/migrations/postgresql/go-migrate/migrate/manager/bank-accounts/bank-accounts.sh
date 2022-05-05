#!/usr/bin/env bash

DESTINATION="internal/adaptors/datastore/postgresql/manager/bank-accounts/migrations"
SCHEMA_NAME="create_bank_accounts_models"

migrate create -ext sql -dir $DESTINATION -seq $SCHEMA_NAME
