lint:
	@echo "linting..."
	@golangci-lint run ./...

test:
	@echo "testing..."
	@go test -v ./...

build-server: lint test
	@go build cmd/rest-server/main.go -o bin/rest-server