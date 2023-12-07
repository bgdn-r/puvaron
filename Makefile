gocmd := go
gobuild := $(gocmd) build
gotest := $(gocmd) test
goclean := $(gocmd) clean
gofmt := $(gocmd) fmt
govendor := $(gocmd) mod vendor
lintcmd := golangci-lint
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
	$(gobuild) -o $(bin_name) -v *.go

test:
	$(gotest) -v ./...

clean:
	$(goclean)
	rm -f $(bin_name)

run: build
	./$(bin_name)

lint:
	$(golint) -v ./...

fmt:
	$(gofmt) ./...

vendor:
	$(govendor)

.PHONY: all build test clean run lint fmt docker-build docker-compose-up docker-compose-down
