#!/usr/bin/env bash

# RELATIVE_PATH_FROM_ROOT=../../../.. # => if called from the file
RELATIVE_PATH_FROM_ROOT=. # if called from the Makefile

${RELATIVE_PATH_FROM_ROOT}/scripts/migrations/postgresql/sqlc/migrate/manager/devices/devices.sh
