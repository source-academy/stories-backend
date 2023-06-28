OUTPUT_PATH=./app

.PHONY: dev build lint

dev:
	@GO_ENV=development go run main.go
build:
	@go build -o ${OUTPUT_PATH} main.go
lint:
	@golangci-lint run
