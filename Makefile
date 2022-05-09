# Makefile

sqlc-gen:
	./scripts/codegen/sqlc/sqlc-generate.sh
	go mod tidy
	go mod vendor

mock-generate:
	go get -d github.com/golang/mock/mockgen
	go mod vendor
	go generate ./...
	go mod vendor

get-migrator:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go mod tidy
	go mod vendor

migrate-init:
	./scripts/migrations/postgresql/go-migrate/run/types/init/run.sh

migrate:
	@./scripts/migrations/postgresql/go-migrate/run/migrate.sh ${DOMAIN_NAME} ${DB_ENV_NAME} ${TYPE} ${OFFSET}

migrate-all-up:
	@./scripts/migrations/postgresql/go-migrate/run/types/up/migrate-all-up.sh

migrate-all-down:
	@./scripts/migrations/postgresql/go-migrate/run/types/down/migrate-all-down.sh

populate-playground:
	echo @./scripts/migrations/postgresql/go-migrate/run/

test-unit:
	go test -race -v ./...

test-unit-cover:
	go test -coverprofile ./docs/reports/tests/unit/cover.out ./...

test-unit-cover-report:
	go tool cover -html=./docs/reports/tests/unit/cover.out

build-local:
	./scripts/build/build-local.sh

clean-local-build:
	rm -rf build
	build-local
