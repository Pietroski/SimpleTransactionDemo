#!/usr/bin/env bash

sqlc init -f internal/adaptors/datastore/postgresql/auth/sqlc/config/auth.sqlc.yaml
sqlc init -f internal/adaptors/datastore/postgresql/manager/devices/sqlc/config/devices.sqlc.yaml
sqlc init -f internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc/config/bank-accounts.sqlc.yaml
