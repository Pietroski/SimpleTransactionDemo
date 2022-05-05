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
#	go get -u -d github.com/golang-migrate/migrate/cmd/migrate
#	cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
#	git checkout $TAG  # e.g. v4.15.2
#	go build -tags 'postgres' -ldflags="-X main.Version=$(git describe --tags)" -o $GOPATH/bin/migrate $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
#	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
#	go mod tidy
#	go mod vendor

migrate-upgrade:
	./scripts/migrations/postgresql/go-migrate/upgrade/run.sh

migrate-all-up:
	@./scripts/migrations/postgresql/go-migrate/run/migrate-all-up.sh

build-local:
	./scripts/build/build-local.sh

clean-local-build:
	rm -rf build
	build-local
