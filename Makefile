include .env
export $(shell sed 's/=.*//' .env)

pretty:
	gofmt -s -w .

build:
	go build cmd/main.go

run: build
	./main

test:
	go test `go list ./... | grep -v cmd`
