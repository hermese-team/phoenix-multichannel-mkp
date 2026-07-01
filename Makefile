.PHONY: run-server run-worker auth-bot build test lint up down

## Local dev — fill in real values or export env vars before running
PARTNER_ID  ?= 1235993
PARTNER_KEY ?= shpk746f486656425876677846657a577650567a726774414a7a616149794a42
SHOP_ID     ?= 225997847
ACCESS_TOKEN ?=
REFRESH_TOKEN ?=

ANTI_BOT_COOKIE ?=
BASE_URL ?= https://openplatform.sandbox.test-stable.shopee.sg
AUTH_CODE ?=
SPC_SI ?=
SC_DFP ?=
SPC_SEC_SI ?=

auth-bot:
	SHOPEE_SELLER_USERNAME=$(SELLER_USERNAME) \
	SHOPEE_SELLER_PASSWORD=$(SELLER_PASSWORD) \
	SHOPEE_PARTNER_ID=$(PARTNER_ID) \
	SHOPEE_PARTNER_KEY=$(PARTNER_KEY) \
	SHOPEE_SHOP_ID=$(SHOP_ID) \
	SHOPEE_BASE_URL=$(BASE_URL) \
	SHOPEE_ANTI_BOT_COOKIE=$(ANTI_BOT_COOKIE) \
	SHOPEE_AUTH_CODE=$(AUTH_CODE) \
	SHOPEE_SPC_SI=$(SPC_SI) \
	SHOPEE_SC_DFP=$(SC_DFP) \
	SHOPEE_SPC_SEC_SI=$(SPC_SEC_SI) \
	go run ./cmd/authbot

run-server:
	SHOPEE_PARTNER_ID=$(PARTNER_ID) \
	SHOPEE_PARTNER_KEY=$(PARTNER_KEY) \
	SHOPEE_SHOP_ID=$(SHOP_ID) \
	SHOPEE_ACCESS_TOKEN=$(ACCESS_TOKEN) \
	SHOPEE_REFRESH_TOKEN=$(REFRESH_TOKEN) \
	SHOPEE_BASE_URL=https://openplatform.sandbox.test-stable.shopee.sg \
	EXEC=server go run ./cmd/server

run-worker:
	SHOPEE_PARTNER_ID=$(PARTNER_ID) \
	SHOPEE_PARTNER_KEY=$(PARTNER_KEY) \
	SHOPEE_SHOP_ID=$(SHOP_ID) \
	SHOPEE_ACCESS_TOKEN=$(ACCESS_TOKEN) \
	SHOPEE_REFRESH_TOKEN=$(REFRESH_TOKEN) \
	SHOPEE_BASE_URL=https://openplatform.sandbox.test-stable.shopee.sg \
	go run ./cmd/worker

build:
	go build -o ./build/server ./cmd/server
	go build -o ./build/worker ./cmd/worker

test:
	go test ./... -race -count=1

lint:
	golangci-lint run ./...

up:
	docker compose -f deployments/docker-compose.yml up -d

down:
	docker compose -f deployments/docker-compose.yml down

