-include .env

all: build

bin:
	mkdir -p bin

build: bin
	go build -o bin ./cmd/...

test:
	go test -v ./...
