ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
.DEFAULT_GOAL := help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

swag: ## swag init
	swag init -g internal/delivery/http/route/route.go
.PHONY: swag

build: ## build binary file
	GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o ./bin/app ./cmd
.PHONY: build

run: swag ## swag then run
	go mod tidy && go mod download && \
	GIN_MODE=debug go run -tags migrate cmd/main.go
.PHONY: run

setup:
	go install github.com/swaggo/swag/cmd/swag@v1.8.12
