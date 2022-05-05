#!/usr/bin/env bash

# This is a test file for bash script testing commands

ARG=$1

# TEST=$(./def-src-url.run)
TEST=$(./scripts/migrations/postgresql/go-migrate/run/def-src-url.sh) # called from the Makefile

echo "this is a test -> $TEST"

declare -A migrations=(
       [test]="asdffdsa"
       [test-again]="fdsafdsa"
)

for key in "${!migrations[@]}" ; do
    echo "$key -> ${migrations[$key]}"
done

for key in "${migrations[@]}" ; do
    echo "$key -> ${migrations[$key]}"
done

declare -A assArray1=(
       ["test"]="asdffdsa"
)
for key in "${!assArray1[@]}"; do echo "$key => ${assArray1[$key]}"; done

echo "${migrations[$ARG]}"
