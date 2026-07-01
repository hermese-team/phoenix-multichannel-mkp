## Phoenix Multichannel Marketplace — Monorepo root Makefile
##
## Per-service tasks delegate to the service's own Makefile.
## Run `make help` to list available targets.

.DEFAULT_GOAL := help

SERVICES := \
  services/order-ingestion \
  services/channel-adapter-shopee

LIBRARIES := \
  libraries/go/telemetry \
  libraries/go/eventing \
  libraries/go/postgres \
  libraries/go/channel-sdk \
  libraries/go/idempotency

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2}'

# ── Workspace ─────────────────────────────────────────────────────────────────

.PHONY: tidy
tidy: ## go mod tidy for all services and libraries
	@for dir in $(SERVICES) $(LIBRARIES); do \
	  echo "→ tidy $$dir"; \
	  (cd $$dir && go mod tidy) || exit 1; \
	done

.PHONY: build
build: ## Build all services
	@for dir in $(SERVICES); do \
	  echo "→ build $$dir"; \
	  (cd $$dir && go build ./...) || exit 1; \
	done

.PHONY: test
test: ## Run tests for all services and libraries
	@for dir in $(SERVICES) $(LIBRARIES); do \
	  echo "→ test $$dir"; \
	  (cd $$dir && go test ./... -race -count=1) || exit 1; \
	done

.PHONY: lint
lint: ## Run golangci-lint for all services and libraries
	@for dir in $(SERVICES) $(LIBRARIES); do \
	  echo "→ lint $$dir"; \
	  (cd $$dir && golangci-lint run ./...) || exit 1; \
	done

# ── Per-service shortcuts ──────────────────────────────────────────────────────

.PHONY: run-order-ingestion
run-order-ingestion: ## Run order-ingestion service
	$(MAKE) -C services/order-ingestion run

.PHONY: run-channel-adapter-shopee
run-channel-adapter-shopee: ## Run channel-adapter-shopee service
	$(MAKE) -C services/channel-adapter-shopee run-server
