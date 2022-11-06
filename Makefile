.DEFAULT_GOAL := lint

NAME := $(shell basename $(CURDIR))

.PHONY: clean
clean: clean-mock
	@echo "Cleaning ${NAME} ..."
	@go clean -i ./...
	@rm -rf bin

.PHONY: all
all: clean lint build test clean-mock

.PHONY: build
build: clean
	@echo "Building ${NAME}..."
	@go build -tags '$(TAGS)' ./...
	@go build -tags '$(TAGS)' -o ./examples/bin/simple/watch ./examples/simple/watch
	@go build -tags '$(TAGS)' -o ./examples/bin/simple/throttle ./examples/simple/throttle

.PHONY: test
test: build mock
	@echo "Testing ${NAME}..."
	@go test -tags '$(TAGS)' ./... -cover -race -v

.PHONY: lint
lint: mock
	@echo "Linting ${NAME}..."
	@go vet -tags '$(TAGS)' ./...

	@golangci-lint run #https://golangci-lint.run/usage/install/

.PHONY: format
format:
	@echo "Formatting ${NAME}..."
	@go mod tidy
	@gofumpt -l -w . #go install mvdan.cc/gofumpt@latest

.PHONY: mock
mock:
	@echo "Generating mocks"
	go install github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen@v1.6.0
	mkdir -p mock

.PHONY: clean-mock
clean-mock:
	@echo "Cleaning mocks ..."
	@rm -rf examples/bin/*
	@rm -rf bin/*
	@rm -rf mock/*