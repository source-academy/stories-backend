PROGRAM_ENTRYPOINT=./main.go
OUTPUT_PATH=./app

.PHONY: dev

dev:
	@GO_ENV=development go run ${PROGRAM_ENTRYPOINT}
build:
	@go build -o ${OUTPUT_PATH} ${PROGRAM_ENTRYPOINT}
