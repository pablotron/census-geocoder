.PHONY=all test

all:
	go build

test:
	go test -short ./...

fulltest:
	go test ./...

lint:
	go vet ./... && staticcheck ./... && golangci-lint run ./...
