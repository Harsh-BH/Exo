# Makefile for the EXO CLI — https://github.com/Harsh-BH/Exo
APP_NAME  := exo
MODULE    := github.com/Harsh-BH/Exo
GOFLAGS   := -ldflags="-s -w -X $(MODULE)/cmd/exo.version=$(shell git describe --tags --always --dirty 2>/dev/null || echo dev)"
PLATFORMS := linux/amd64 darwin/amd64 darwin/arm64 windows/amd64

.PHONY: all build install test coverage lint fmt vet clean release docs help

all: build

## build: Compile the exo binary into ./bin/exo
build:
	@echo "→ Building bin/$(APP_NAME) …"
	@mkdir -p bin
	go build $(GOFLAGS) -o bin/$(APP_NAME) .

## install: Install exo to $GOPATH/bin (or ~/go/bin)
install:
	@echo "→ Installing $(APP_NAME) …"
	go install $(GOFLAGS) .

## test: Run all unit tests with the race detector
test:
	@echo "→ Running tests …"
	go test -v -race ./...

## coverage: Run tests and open HTML coverage report
coverage:
	@echo "→ Generating coverage report …"
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

## fmt: Format all Go source files
fmt:
	@echo "→ Formatting …"
	gofmt -s -w .

## vet: Run go vet
vet:
	@echo "→ Vetting …"
	go vet ./...

## lint: Run golangci-lint (install: https://golangci-lint.run/usage/install/)
lint:
	@echo "→ Linting …"
	golangci-lint run ./...

## release: Cross-compile binaries for all platforms into ./bin/
release:
	@echo "→ Cross-compiling …"
	@mkdir -p bin
	GOOS=linux   GOARCH=amd64 go build $(GOFLAGS) -o bin/$(APP_NAME)-linux-amd64   .
	GOOS=darwin  GOARCH=amd64 go build $(GOFLAGS) -o bin/$(APP_NAME)-darwin-amd64  .
	GOOS=darwin  GOARCH=arm64 go build $(GOFLAGS) -o bin/$(APP_NAME)-darwin-arm64  .
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -o bin/$(APP_NAME)-windows-amd64.exe .
	@echo "→ Binaries in ./bin/"

## docs: Regenerate CLI documentation into docs/cli/
docs:
	@echo "→ Generating CLI docs …"
	go run . docs

## clean: Remove build artefacts
clean:
	@echo "→ Cleaning …"
	rm -rf bin/$(APP_NAME) coverage.out

## help: Show this help
help:
	@grep -E '^## [a-z]' Makefile | sed 's/## /  /'
