GO_FILES = $(shell find src/*.go)

run:
	go run $(GO_FILES)
build:
	go build -o dist/main $(GO_FILES)