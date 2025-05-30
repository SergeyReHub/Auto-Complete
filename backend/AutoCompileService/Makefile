GO_OUT_DIR := ./pkg
PROTO_DIR := ./api/auto-complite
GO_OUT_DIR = ./pkg/proto
PROTO_FILES := $(shell find $(PROTO_ROOT) -name '*.proto')

PHONY: all generate build run

all: build run

# Генерация кода для gRPC
generate:
	@echo "Generating code from proto files in $(PROTO_DIR)..."
	protoc -I $(PROTO_DIR) -I . \
		--go_out=$(GO_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

# Сборка бинарника
build:
	@echo "Building binary..."
	go build -o bin/server ./cmd/service_auto_complite
	@echo "Binary built."

# Запуск сервера
run:
	@echo "Starting server..."
	./bin/server
