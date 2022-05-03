#!/usr/bin/env bash

DESTINATION="internal/adaptors/datastore/postgresql/sqlc/manager/devices/migrations"
SCHEMA_NAME="create_devices_models"

migrate create -ext sql -dir $DESTINATION -seq $SCHEMA_NAME
