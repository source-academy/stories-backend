PROGRAM_ENTRYPOINT=./main.go
OUTPUT_PATH=./app

DB_SCRIPT_ENTRYPOINT=./scripts/db.go
DB_TARGETS=db_migrate db_rollback db_status

.PHONY: dev build test testCI coverage lint format hooks $(DB_TARGETS)

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

db_migrate: db_status
	@GO_ENV=development go run ${DB_SCRIPT_ENTRYPOINT} migrate $(steps)
db_rollback:
	@GO_ENV=development go run ${DB_SCRIPT_ENTRYPOINT} rollback $(steps)
	@$(MAKE) db_status
db_status:
	@GO_ENV=development go run ${DB_SCRIPT_ENTRYPOINT} status