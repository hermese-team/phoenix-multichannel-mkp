## order-ingestion service Makefile
## Run from repo root: make -C services/order-ingestion <target>

.DEFAULT_GOAL := help

# ── Tools ────────────────────────────────────────────────────────────────────
SQLC    ?= $(shell which sqlc)
MIGRATE ?= $(shell which migrate)
LINT    ?= golangci-lint

# ── Migration config ─────────────────────────────────────────────────────────
MIGRATIONS_DIR ?= file://migrations
DB_URL         ?= pgx5://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSLMODE)

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2}'

# ── Build ─────────────────────────────────────────────────────────────────────

.PHONY: build
build: ## Build the API binary
	go build -o ../../bin/order-ingestion ./cmd/api

.PHONY: run
run: ## Run the API server locally
	go run ./cmd/api

# ── Test / Lint ───────────────────────────────────────────────────────────────

.PHONY: test
test: ## Run tests
	go test ./... -race -count=1

.PHONY: lint
lint: ## Run golangci-lint
	$(LINT) run ./...

.PHONY: tidy
tidy: ## go mod tidy
	go mod tidy

# ── Code generation ───────────────────────────────────────────────────────────

.PHONY: sqlc
sqlc: ## Generate type-safe DB code via sqlc
	$(SQLC) generate

# ── Migrations ────────────────────────────────────────────────────────────────

.PHONY: migrate-up
migrate-up: ## Apply all pending migrations
	$(MIGRATE) -path migrations -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down: ## Roll back the last migration
	$(MIGRATE) -path migrations -database "$(DB_URL)" down 1

.PHONY: migrate-create
migrate-create: ## Create a new migration: make migrate-create NAME=add_foo
	@test -n "$(NAME)" || (echo "Usage: make migrate-create NAME=<migration_name>" && exit 1)
	$(MIGRATE) create -ext sql -dir migrations -seq $(NAME)

# ── Tools install ─────────────────────────────────────────────────────────────

.PHONY: tools
tools: ## Install dev tools (sqlc, migrate, golangci-lint)
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
