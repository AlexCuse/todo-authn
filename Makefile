generate-db:
	go get -modfile codegen.mod github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	cd db && go run -modfile ../codegen.mod github.com/sqlc-dev/sqlc/cmd/sqlc generate

generate-api:
	go get -modfile codegen.mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	cd api && go run -modfile ../codegen.mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=oapi-codegen.yml api.yml

generate: generate-db generate-api

build-ui:
	cd ui && npm run build

lint:
	@if which golangci-lint >/dev/null ; then golangci-lint run --config .golangci.yml ; else echo "WARNING: golangci-lint not installed."; fi

run:
	docker-compose up --build

ssl:
	mkcert todoauthn.com "*.todoauthn.com" localhost
	mv todoauthn.com* .certs/
	mkcert -install

all: generate build-ui run