# Variables
APP_NAME=simplerestapi

# Tasks
default: run

run:
	@echo "Running server..."
	@go run internal/main.go