dir := $(shell mktemp -d)

.PHONY: install-swagger
install-swagger:
	git clone https://github.com/go-swagger/go-swagger "$(dir)"; \
	cd "$(dir)" && git checkout tags/v0.27.0 && go install -ldflags "-X github.com/go-swagger/go-swagger/cmd/swagger/commands.Version=$(git describe --tags) -X github.com/go-swagger/go-swagger/cmd/swagger/commands.Commit=$(git rev-parse HEAD)" ./cmd/swagger
	rm -rf "$(dir)"

.PHONY: init-swagger
init-swagger:
	swagger init spec \
      --title "Flash application" \
      --description "To send money between users" \
      --version 0.0.1 \
      --scheme http \
      --consumes application/json \
      --produces application/json

.PHONY: validate
validate:
	swagger validate ./docs/swagger.yml

.PHONY: gen
gen:
	swagger generate server -t api -A flash -P models.Principle -f ./docs/swagger.yml --exclude-main

.PHONY: run
run:
	go run cmd/flash-server/main.go --port 8000

.PHONY: build
build:
	go build -o bin/flash-server cmd/flash-server/main.go


.PHONY: docker-up
docker-up:
	docker-compose up -d

.PHONY: sql-init
sql-init:
	docker exec -i node1 sqlcmd < internal/database/schema.sql



