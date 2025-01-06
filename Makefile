include .env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN=$(PG_DSN)


run:
	docker-compose up -d
	go run cmd/grpc_server/main.go

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	make install-golangci-lint

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml

generate:
	make generate-message-api

generate-message-api:
	mkdir -p pkg/message_v1
	protoc --proto_path api/message_v1 \
	--go_out=pkg/message_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/message_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/message_v1/message.proto


docker-build:
	docker buildx build --no-cache .

local-migration-up:
	bin/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

grpc-err-load-test:
	ghz \
	--proto api/message_v1/message.proto \
	--call message_v1.ChatV1/CreateChat \
	--data '{"created_by": "62180", "name": "fugiat do"}' \
	--rps 1 \
	--total 10 \
	--insecure \
	localhost:50053s