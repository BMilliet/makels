.PHONY: help build run fmt deps test clean install

# Variables
BINARY_NAME=makels
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")
GO_PACKAGES=$(shell go list ./... | grep -v /vendor/)

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

deps: ## Download and install dependencies
	@echo "📦 Installing dependencies..."
	go mod tidy
	go mod download

build: ## Build the application
	@echo "🏗️  Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) .
	mkdir -p ~/.makels
	mv ${BINARY_NAME} ~/.makels/${BINARY_NAME}
	@echo "✓ Binary installed at ~/.makels/$(BINARY_NAME)"

install: ## Add ~/.makels to PATH (show instructions)
	@echo "To use makels globally, add this to your shell config (~/.zshrc or ~/.bashrc):"
	@echo ""
	@echo "  export PATH=\"\$$HOME/.makels:\$$PATH\""
	@echo ""
	@echo "Then run: source ~/.zshrc  (or source ~/.bashrc)"

run: ## Run the application without building binary
	@echo "🚀 Running $(BINARY_NAME)..."
	go run main.go

fmt: ## Format Go code
	@echo "✨ Formatting code..."
	gofmt -s -w $(GO_FILES)
	go fmt $(GO_PACKAGES)

test: ## Run tests
	@echo "🧪 Running tests..."
	go test -v ./...

vet: ## Run go vet
	@echo "🔍 Running go vet..."
	go vet $(GO_PACKAGES)

lint: fmt vet ## Run linters (fmt + vet)
	@echo "✅ Linting complete"

clean: ## Remove binary and build artifacts
	@echo "🧹 Cleaning..."
	rm -f $(BINARY_NAME)
	@echo "✓ Clean complete"
