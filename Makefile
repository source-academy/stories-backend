OUTPUT_PATH=./app

.PHONY: dev

dev:
	@go run main.go
build:
	@go build -o ${OUTPUT_PATH} main.go
