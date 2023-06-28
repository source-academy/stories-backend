PROGRAM_ENTRYPOINT=./main.go
OUTPUT_PATH=./app

.PHONY: dev build test coverage lint format

dev:
	@GO_ENV=development go run ${PROGRAM_ENTRYPOINT}
build:
	@go build -o ${OUTPUT_PATH} ${PROGRAM_ENTRYPOINT}
test:
	@go test -v ./...
coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
lint:
	@golangci-lint run
format:
	@go fmt ./...
