.PHONY: build test lint install clean

build:
	go build -o bin/keen ./cmd/jira-pp-cli-pp-cli

test:
	go test ./...

lint:
	golangci-lint run

install:
	go install ./cmd/jira-pp-cli-pp-cli

clean:
	rm -rf bin/

build-mcp:
	go build -o bin/keen-mcp ./cmd/jira-pp-cli-pp-mcp

install-mcp:
	go install ./cmd/jira-pp-cli-pp-mcp

build-all: build build-mcp
