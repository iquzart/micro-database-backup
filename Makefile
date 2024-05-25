.DEFAULT_GOAL := help
.PHONY: help clean build run docker-build docker-run cleanup

# Define variables
BINARY_NAME := micro-db-backup
MAIN_FILE := cmd/main.go
LDFLAGS := -ldflags="-s -w"
DOCKER_IMAGE := micro-db-backup
DOCKER_CONTAINER := micro-db-backup
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

help: ## Display this help message
	@echo "Usage: make <command>"
	@echo ""
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

clean: ## Remove binary file and clean up generated resources
	rm -f $(BINARY_NAME)
	docker rm -f $(DOCKER_CONTAINER) 2>/dev/null || true
	docker rmi -f $(DOCKER_IMAGE) 2>/dev/null || true

build: ## Build the binary file
	go build $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_FILE)

run: ## Build and run the binary file
	go run $(LDFLAGS) $(MAIN_FILE)

deps: ## Install go modules dependencies
	go mod tidy

lint: ## Run go lint on all files
	golangci-lint run $(GO_FILES)

test: ## Run tests
	go test ./...

fmt: ## Format the code
	go fmt $(GO_FILES)

vet: ## Run go vet
	go vet $(GO_FILES)

docker-build: ## Build the Docker image
	docker build -t $(DOCKER_IMAGE) -f Containerfile .

docker-run: docker-build ## Run the Docker container
	docker run --name $(DOCKER_CONTAINER) -d \
	  -e MONGODB_HOST="localhost" \
	  -e MONGODB_PORT="27017" \
	  -e MONGODB_USERNAME="admin" \
	  -e MONGODB_PASSWORD="password" \
	  -e BACKUP_DIR="/tmp/micro-database-backup/mongodb" \
	  -e DATABASE_NAME="mydb" \
	  $(DOCKER_IMAGE)

cleanup: ## Cleanup Docker resources
	docker rm -f $(DOCKER_CONTAINER) 2>/dev/null || true
	docker rmi -f $(DOCKER_IMAGE) 2>/dev/null || true
	docker system prune -f

