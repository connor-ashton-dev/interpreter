# Makefile for crafting_interpreters

BINARY=interpreter

.PHONY: all build run test lint clean

all: build

build:
	go build -o $(BINARY) main.go

run: build
	./$(BINARY)

test:
	go test ./...

lint:
	golangci-lint run ./...

clean:
	rm -f $(BINARY) 