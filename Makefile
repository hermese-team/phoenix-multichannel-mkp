APP_NAME    := phoenix-multichannel-mkp
MODULE      := github.com/ascend/phoenix-multichannel-mkp
CMD_PATH    := ./cmd/api
BIN_DIR     := bin
BINARY      := $(BIN_DIR)/$(APP_NAME)

IMAGE_NAME  := $(APP_NAME)
IMAGE_TAG   ?= latest

GO          := go
GOFLAGS     ?=
MOCKGEN     := $(GO) tool mockgen
SWAG        := $(GO) tool swag

# Swagger / OpenAPI
SWAGGER_DIR  := docs
SWAGGER_MAIN := $(CMD_PATH)/main.go

# Inject build metadata
VERSION     ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT      ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_TIME  := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS     := -s -w \
               -X $(MODULE)/pkg/build.Version=$(VERSION) \
               -X $(MODULE)/pkg/build.Commit=$(COMMIT) \
               -X $(MODULE)/pkg/build.BuildTime=$(BUILD_TIME)

.DEFAULT_GOAL := help

# ── Help ──────────────────────────────────────────────────────────────────────

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# ── Build ─────────────────────────────────────────────────────────────────────

.PHONY: build
build: ## Build binary to bin/
	@mkdir -p $(BIN_DIR)
	$(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BINARY) $(CMD_PATH)

.PHONY: run
run: ## Run directly via go run
	$(GO) run $(CMD_PATH)

# ── Test ──────────────────────────────────────────────────────────────────────

.PHONY: test
test: ## Run all tests
	$(GO) test ./... -race -count=1

.PHONY: test-cover
test-cover: ## Run tests and show coverage
	$(GO) test ./... -race -coverprofile=coverage.out -covermode=atomic
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# ── Codegen ───────────────────────────────────────────────────────────────────

.PHONY: tools
tools: ## Register codegen tools in go.mod (run once; needs network)
	$(GO) get -tool go.uber.org/mock/mockgen@latest
	$(GO) get -tool github.com/swaggo/swag/cmd/swag@latest

.PHONY: generate
generate: ## Generate mocks for repository/service interfaces
	@mkdir -p internal/repository/mocks internal/service/mocks
	$(MOCKGEN) -source=internal/repository/repository.go -destination=internal/repository/mocks/mock_repository.go -package=mocks
	$(MOCKGEN) -source=internal/service/service.go -destination=internal/service/mocks/mock_service.go -package=mocks

.PHONY: swagger
swagger: ## Generate Swagger/OpenAPI docs from annotations
	$(SWAG) init -g $(SWAGGER_MAIN) -o $(SWAGGER_DIR) --parseInternal

.PHONY: swagger-fmt
swagger-fmt: ## Format Swagger annotation comments
	$(SWAG) fmt

# ── Quality ───────────────────────────────────────────────────────────────────

.PHONY: lint
lint: ## Run golangci-lint (must be installed)
	golangci-lint run ./...

.PHONY: vet
vet: ## Run go vet
	$(GO) vet ./...

.PHONY: fmt
fmt: ## Format all Go source files
	$(GO) fmt ./...

.PHONY: tidy
tidy: ## Tidy and verify go.mod / go.sum
	$(GO) mod tidy
	$(GO) mod verify

.PHONY: check
check: fmt vet test ## fmt + vet + test

# ── Docker ────────────────────────────────────────────────────────────────────

.PHONY: docker-build
docker-build: ## Build Docker image
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg BUILD_TIME=$(BUILD_TIME) \
		-t $(IMAGE_NAME):$(IMAGE_TAG) \
		.

.PHONY: docker-run
docker-run: ## Run container (PORT default 8080)
	docker run --rm \
		-p ${PORT:-8080}:8080 \
		-e GIN_MODE=release \
		--name $(APP_NAME) \
		$(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: docker-stop
docker-stop: ## Stop running container
	docker stop $(APP_NAME) 2>/dev/null || true

# ── Cleanup ───────────────────────────────────────────────────────────────────

.PHONY: clean
clean: ## Remove build artifacts and coverage files
	rm -rf $(BIN_DIR) coverage.out coverage.html $(SWAGGER_DIR)
