PROGRAM_ENTRYPOINT=./main.go
OUTPUT_PATH=./app

.PHONY: dev build test coverage lint format hooks

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
hooks:
	@cp -f ./hooks/pre-commit.sh ./.git/hooks/pre-commit
	@chmod 755 ./.git/hooks/pre-commit
	@cp -f ./hooks/pre-push.sh ./.git/hooks/pre-push
	@chmod 755 ./.git/hooks/pre-push
	@echo "Hooks installed successfully!"
