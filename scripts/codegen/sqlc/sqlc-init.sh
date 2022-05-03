#!/usr/bin/env bash

sqlc init -f internal/adaptors/datastore/postgresql/sqlc/auth/auth.sqlc.yaml
sqlc init -f internal/adaptors/datastore/postgresql/sqlc/manager/devices/devices.sqlc.yaml
sqlc init -f internal/adaptors/datastore/postgresql/sqlc/manager/bank-accounts/bank-accounts.sqlc.yaml
