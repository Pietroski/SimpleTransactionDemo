# Makefile

mock-generate:
	go get -d github.com/golang/mock/mockgen
	go mod vendor
	go generate ./...
	go mod vendor

build-local:
	./scripts/build/build-local.sh

migrate:
	./scripts/migrations/postgresql/sqlc/run/run.sh

migrate-all:
	./scripts/migrations/run/migrate-all.sh

sqlc-gen:
	./scripts/codegen/sqlc/sqlc-generate.sh
	go mod tidy
	go mod vendor

clean-build:
	rm -rf build
	build
