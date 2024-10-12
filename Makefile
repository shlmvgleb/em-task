-include .env

all: build

bin:
	mkdir -p bin

docs:
	swag init -g ./cmd/app/main.go -o cmd/docs

build: bin docs
	go build -o bin ./cmd/...

test:
	go test -v ./...
