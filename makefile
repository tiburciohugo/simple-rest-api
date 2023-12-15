# Variables
APP_NAME=simplerestapi
GOFLAGS := -mod=mod

default: run

# Tasks
build:
	@echo Building Go application...
	@go build $(GOFLAGS) -o bin/$(APP_NAME) main.go

run:
	@echo Running Go application...
	@go run $(GOFLAGS) main.go

docs:
	@echo Generating Swagger documentation...
	@swag init

clean:
	@rm -rf bin/$(APP_NAME) doc/index.html

.PHONY: build run docs clean