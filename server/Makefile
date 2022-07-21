.PHONY: init build detector test

init:
		go mod tidy
build:
		make init
		make test
		go build --race -o build/server -v ./cmd/app
detector:
		make build
		go run ./cmd/app
test:
		make init
		go test -v ./internal/app/counter
		go test -v ./internal/app/server
		golint ./cmd/app
		go vet ./cmd/app


.DEFAULT_GOAL := detector
