SHELL := /bin/bash
GO    := go

.DEFAULT_GOAL := help

# ── Help ──────────────────────────────────────────────────────────────────────

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2}'

# ── Build & Run ───────────────────────────────────────────────────────────────

.PHONY: build
build: ## Build all sell-channel binaries to bin/
	$(GO) build -o bin/shopee-server    ./sellchannel/shopee/cmd/server
	$(GO) build -o bin/shopee-consumer  ./sellchannel/shopee/cmd/consumer
	$(GO) build -o bin/shopee-scheduler ./sellchannel/shopee/cmd/scheduler
	$(GO) build -o bin/lazada-server    ./sellchannel/lazada/cmd/server
	$(GO) build -o bin/lazada-order-scheduler ./sellchannel/lazada/cmd/order-scheduler
	$(GO) build -o bin/lazada-consumer  ./sellchannel/lazada/cmd/consumer
	$(GO) build -o bin/lazada-product-scheduler ./sellchannel/lazada/cmd/product-scheduler
	$(GO) build -o bin/lazada-product-create-scheduler ./sellchannel/lazada/cmd/product-create-scheduler
	$(GO) build -o bin/lazada-product-update-scheduler ./sellchannel/lazada/cmd/product-update-scheduler
	$(GO) build -o bin/lazada-product-offsale-scheduler ./sellchannel/lazada/cmd/product-offsale-scheduler
	$(GO) build -o bin/lazada-sellable-stock-scheduler ./sellchannel/lazada/cmd/sellable-stock-scheduler
	$(GO) build -o bin/lazada-fulfillment-scheduler ./sellchannel/lazada/cmd/fulfillment-scheduler
	$(GO) build -o bin/lazada-token     ./sellchannel/lazada/cmd/lazada-token

.PHONY: check-consumers
check-consumers: ## List running consumer processes (does not kill)
	@pgrep -fl 'exe/consumer' || pgrep -fl 'bin/[a-z]*-consumer' || echo "no consumer processes"

.PHONY: kill-consumers
kill-consumers: ## Kill orphaned consumer processes left behind by `go run`
	@pkill -f 'exe/consumer'        || true
	@pkill -f 'bin/[a-z]*-consumer' || true
	@echo "killed orphaned consumers (if any)"

.PHONY: run-lazada-server
run-lazada-server: ## Run the lazada server
	$(GO) run ./sellchannel/lazada/cmd/server

.PHONY: run-lazada-order-scheduler
run-lazada-order-scheduler: ## Run the lazada order-sync scheduler
	$(GO) run ./sellchannel/lazada/cmd/order-scheduler

.PHONY: run-lazada-consumer
run-lazada-consumer: ## Run the lazada webhook->order consumer
	$(GO) run ./sellchannel/lazada/cmd/consumer

.PHONY: run-lazada-product-scheduler
run-lazada-product-scheduler: ## Run the lazada product price/stock sync scheduler
	$(GO) run ./sellchannel/lazada/cmd/product-scheduler

.PHONY: run-lazada-product-create-scheduler
run-lazada-product-create-scheduler: ## Run the lazada product-create scheduler
	$(GO) run ./sellchannel/lazada/cmd/product-create-scheduler

.PHONY: run-lazada-product-update-scheduler
run-lazada-product-update-scheduler: ## Run the lazada product-update scheduler
	$(GO) run ./sellchannel/lazada/cmd/product-update-scheduler

.PHONY: run-lazada-product-offsale-scheduler
run-lazada-product-offsale-scheduler: ## Run the lazada off-sale scheduler (deactivate/remove)
	$(GO) run ./sellchannel/lazada/cmd/product-offsale-scheduler

.PHONY: run-lazada-sellable-stock-scheduler
run-lazada-sellable-stock-scheduler: ## Run the lazada sellable-stock scheduler (adjust/update)
	$(GO) run ./sellchannel/lazada/cmd/sellable-stock-scheduler

.PHONY: run-lazada-category
run-lazada-category: ## Browse Lazada category tree / attributes (LAZADA_CATEGORY_ID=<id> for attrs)
	$(GO) run ./sellchannel/lazada/cmd/lazada-category

.PHONY: run-lazada-fulfillment-scheduler
run-lazada-fulfillment-scheduler: ## Run the lazada fulfillment scheduler (dropship + own-fleet)
	$(GO) run ./sellchannel/lazada/cmd/fulfillment-scheduler

.PHONY: seed-lazada-token
seed-lazada-token: ## Seed Lazada tokens into Redis (REFRESH=<t> [ACCESS=<t>])
	$(GO) run ./sellchannel/lazada/cmd/lazada-token -refresh="$(REFRESH)" -access="$(ACCESS)"

.PHONY: run-server
run-server: ## Run the shopee server
	$(GO) run ./sellchannel/shopee/cmd/server

.PHONY: run-consumer
run-consumer: ## Run the shopee consumer
	$(GO) run ./sellchannel/shopee/cmd/consumer

.PHONY: run-scheduler
run-scheduler: ## Run the shopee scheduler
	$(GO) run ./sellchannel/shopee/cmd/scheduler

# ── Quality ───────────────────────────────────────────────────────────────────

.PHONY: test
test: ## Run all tests
	$(GO) test ./... -race -count=1

.PHONY: test-cover
test-cover: ## Run tests and open coverage report
	$(GO) test ./... -race -coverprofile=coverage.out -covermode=atomic
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

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
docker-build: ## Build all shopee docker images
	docker build -f sellchannel/shopee/cmd/server/Dockerfile    -t shopee-server    .
	docker build -f sellchannel/shopee/cmd/consumer/Dockerfile  -t shopee-consumer  .
	docker build -f sellchannel/shopee/cmd/scheduler/Dockerfile -t shopee-scheduler .

# ── Cleanup ───────────────────────────────────────────────────────────────────

.PHONY: clean
clean: ## Remove build artifacts and coverage files
	rm -rf bin coverage.out coverage.html
