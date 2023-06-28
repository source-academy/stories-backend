OUTPUT_PATH=./app

.PHONY: dev

dev:
	@GO_ENV=development go run main.go
build:
	@go build -o ${OUTPUT_PATH} main.go
