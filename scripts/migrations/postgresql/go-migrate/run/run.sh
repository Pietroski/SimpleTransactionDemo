#!/usr/bin/env bash

# RELATIVE_PATH_FROM_ROOT=../../../.. # => if called from the file
RELATIVE_PATH_FROM_ROOT=. # if called from the Makefile

${RELATIVE_PATH_FROM_ROOT}/scripts/migrations/postgresql/go-migrate/migrate/auth/auth.sh
${RELATIVE_PATH_FROM_ROOT}/scripts/migrations/postgresql/go-migrate/migrate/manager/devices/devices.sh
${RELATIVE_PATH_FROM_ROOT}/scripts/migrations/postgresql/go-migrate/migrate/manager/bank-accounts/bank-accounts.sh
