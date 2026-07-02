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
	$(GO) build -o bin/lazada-scheduler ./sellchannel/lazada/cmd/scheduler
	$(GO) build -o bin/lazada-consumer  ./sellchannel/lazada/cmd/consumer
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

.PHONY: run-lazada-scheduler
run-lazada-scheduler: ## Run the lazada order-sync scheduler
	$(GO) run ./sellchannel/lazada/cmd/scheduler

.PHONY: run-lazada-consumer
run-lazada-consumer: ## Run the lazada webhook->order consumer
	$(GO) run ./sellchannel/lazada/cmd/consumer

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
