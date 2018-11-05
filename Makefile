.PHONY: build test 

install:
	go install
	pre-commit install

build:
	go build

test:
	go build
	go test ./...
