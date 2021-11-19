# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
BINARY_NAME=bakery
WEBSERVER=./cmd/http

LDFLAGS = -ldflags "-X github.com/cbsinteractive/bakery/handlers.GitSHA=$(shell git rev-parse HEAD)"

all: test build

build: 
	$(GOBUILD) -mod=vendor $(LDFLAGS) -o $(BINARY_NAME) -v $(WEBSERVER)

test: 
	$(GOTEST) -mod=vendor -v -race -count=1 ./...

test_cover:
	GOFLAGS=-p=8 $(GOTEST) -mod=vendor -v -count 1 ./... -race -coverprofile=coverage.txt -covermode=atomic

clean: 
	$(GOCLEAN) -mod=vendor ./...
	rm -f $(BINARY_NAME)

run:
	$(GORUN) $(LDFLAGS) $(WEBSERVER)
