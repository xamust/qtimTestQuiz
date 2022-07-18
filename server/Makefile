.PHONY: init build detector test

init:
		go mod tidy
build:
		make init
		go build -o build/server -v ./cmd/app
detector:
		make build
		go run ./cmd/app
test:
		go test -v ./internal/app/counter
		go test -v ./internal/app/server

.DEFAULT_GOAL := build