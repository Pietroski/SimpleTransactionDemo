#!/usr/bin/env bash

# script echoes a default DB_SOURCE_URL - playground one

DB_DRIVER="postgres"
DB_USERNAME="manager"
DB_PASSWORD="manager"
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="playground_db"
DB_SSL_MODE="disable"
DB_SOURCE_URL="$DB_DRIVER://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$DB_SSL_MODE" # param1=true&param2=false

echo "$DB_SOURCE_URL"
