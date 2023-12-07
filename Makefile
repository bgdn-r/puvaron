gocmd := go
gobuild := $(gocmd) build
gotest := $(gocmd) test
goclean := $(gocmd) clean
gofmt := $(gocmd) fmt
govendor := $(gocmd) mod vendor
lintcmd := golangci-lint

bin_name := puvaron

version := 0.0.1

all: clean fmt lint test run

build:
	$(gobuild) -o $(bin_name) -v cmd/puvaron/main.go

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

.PHONY: all build test clean run lint fmt
