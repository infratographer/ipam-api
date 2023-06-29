all: lint test

PHONY: help all test coverage lint golint clean vendor docker-up docker-down unit-test
GOOS=linux
DB=ipam
DEV_DB=${DB}_dev
TEST_DB=${DB}_test
DEV_URI="postgresql://root@crdb:26257/${DEV_DB}?sslmode=disable"
TEST_URI="postgresql://root@crdb:26257/${TEST_DB}?sslmode=disable"

APP_NAME=ipam-api
PID_FILE=/tmp/${APP_NAME}.pid

help: Makefile ## Print help
	@grep -h "##" $(MAKEFILE_LIST) | grep -v grep | sed -e 's/:.*##/#/' | column -c 2 -t -s#

test: | unit-test

unit-test: ## Runs unit tests
	@echo --- Running unit tests...
	@date --rfc-3339=seconds
	@go test -race -cover -failfast -tags testtools ./...

coverage: ## Generates coverage report
	@echo --- Generating coverage report...
	@date --rfc-3339=seconds
	@go test -race -coverprofile=coverage.out -covermode=atomic -tags testtools -p 1 ./...
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out

lint: golint ## Runs linting

golint:
	@echo --- Running golint...
	@date --rfc-3339=seconds
	@golangci-lint run

clean: ## Clean up all the things
	@echo --- Cleaning...
	@date --rfc-3339=seconds
	@rm -rf ./dist/
	@rm -rf coverage.out
	@go clean -testcache

binary: | generate ## Builds the binary
	@echo --- Building binary...
	@date --rfc-3339=seconds
	@go build -o bin/${APP_NAME} main.go

vendor: ## Vendors dependencies
	@echo --- Downloading dependencies...
	@date --rfc-3339=seconds
	@go mod tidy
	@go mod download

testclient:| background-run .testclient kill-running ## Regenerates the test client in graphclient

.testclient:
	@echo --- Generating test graph client...
	@date --rfc-3339=seconds
	@go generate ./internal/graphclient

dev-nats: ## Initializes nats
	@echo --- Initializing nats
	@date --rfc-3339=seconds
	@.devcontainer/scripts/nats_account.sh

generate: vendor
	@echo --- Generating code...
	@date --rfc-3339=seconds
	@go generate ./...

go-run: ## Runs the app
	@echo --- Running binary...
	@date --rfc-3339=seconds
	@go run main.go serve --playground --dev

background-run:  ## Runs in the app in the background
	@echo --- Running binary in the background...
	@date --rfc-3339=seconds
	@go run main.go serve --pid-file=${PID_FILE} &

kill-running: ## Kills the running binary from pid file
	@echo --- Killing background binary...
	@date --rfc-3339=seconds
	@kill $$(cat ${PID_FILE})
