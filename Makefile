.DEFAULT_GOAL := build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build ./cmd/main.go
