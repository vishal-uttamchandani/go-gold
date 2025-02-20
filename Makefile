lint:
	@echo "linting..."
	@golangci-lint run ./...

test:
	@echo "testing..."
	@go test -v ./...

build-server: lint test
	@go build cmd/rest-server/main.go -o bin/rest-server


run:
	@go run cmd/rest-server/main.go

all: lint test run

.PHONY: all run build-server test lint