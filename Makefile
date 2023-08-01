PROGRAM_ENTRYPOINT=./main.go
OUTPUT_PATH=./app

DB_SCRIPT_ENTRYPOINT=./scripts/db.go ./scripts/create_db.go
DB_TARGETS=db_migrate db_rollback db_status db_create db_drop

.PHONY: dev build test coverage lint format hooks $(DB_TARGETS)

path ?= ./...

dev:
	@GO_ENV=development go run ${PROGRAM_ENTRYPOINT}
build:
	@go build -o ${OUTPUT_PATH} ${PROGRAM_ENTRYPOINT}
test:
	@GO_ENV=test go test -v $(path)
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

$(DB_TARGETS): export GO_ENV=script
db_create:
	@go run ${DB_SCRIPT_ENTRYPOINT} create
db_drop:
	@go run ${DB_SCRIPT_ENTRYPOINT} drop
db_migrate: db_status
	@go run ${DB_SCRIPT_ENTRYPOINT} migrate $(steps)
db_rollback:
	@go run ${DB_SCRIPT_ENTRYPOINT} rollback $(steps)
	@$(MAKE) db_status
db_status:
	@go run ${DB_SCRIPT_ENTRYPOINT} status