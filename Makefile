gocmd := go
golintcmd := golangci-lint
dockercmd := docker

bin_name := puvaron

version := 0.0.1

all: clean fmt lint test build docker-build

docker-build:
	$(dockercmd) build -t $(bin_name) .

docker-compose-up:
	$(dockercmd) compose up -d

docker-compose-down:
	$(dockercmd) compose down

build:
	$(gocmd) build -o $(bin_name) -v cmd/puvaron/main.go

test:
	$(gocmd) test -v ./...

clean:
	$(gocmd) clean
	rm -f $(bin_name)

run: build
	./$(bin_name)

lint:
	$(golintcmd) run ./... -v

fmt:
	$(gocmd) fmt ./...

vendor:
	$(gocmd) mod vendor

tidy:
	$(gocmd) mod tidy

.PHONY: all build test clean run lint fmt docker-build docker-compose-up docker-compose-down vendor tidy
