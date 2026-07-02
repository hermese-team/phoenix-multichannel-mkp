# marketplace-sync

Go monorepo for syncing products, orders, and promotions across sell channels (Shopee, Lazada, TikTok).

## Architecture

Hexagonal Architecture (Ports & Adapters) + DDD. The core rule: **business logic never depends on infrastructure**. Infrastructure depends on the domain, never the other way around.

```
config/                         root config — composes all sub-configs
sellchannel/                    one directory per sell channel
    shopee/
        cmd/                    deployable binaries (one main.go each)
            server/             HTTP server binary
            consumer/           Kafka consumer binary
            scheduler/          Cron scheduler binary
        client/                 Shopee API client (HTTP + signing)
        server/                 Gin HTTP handlers
        consumer/               Kafka message handlers
        scheduler/              Cron job definitions
        config.go               Shopee-specific config struct
        shopee.go               Composition root — wires everything
internal/
    errors.go                   Shared generic errors (ErrNotFound, ErrConflict, ...)
    domain/
        product/
            entity.go           Product struct + constants
            errors.go           Product-specific errors only
            port.go             Repository, MarketplacePublisher, EventPublisher interfaces
        order/
            entity.go
            port.go
    usecase/
        product/
            publish.go          Orchestrates: load → validate → publish → save
        order/
            process_webhook.go
    infrastructure/
        mysql/                  Implements domain repository interfaces
        redis/                  Implements cache interface
        kafka/                  Kafka producer
pkg/
    httpclient/                 Generic HTTP client wrapper
    signer/                     HMAC-SHA256 signing (Shopee API)
```

## Dependency Rules

Read this as "X may import Y":

```
cmd                  → config, sellchannel/shopee
config               → sellchannel/shopee (Config), internal/infrastructure/* (Config)
shopee.go            → internal/infrastructure/*, internal/usecase/*, shopee/client, shopee/server, shopee/consumer, shopee/scheduler
shopee/client        → internal/domain/*, pkg/*
shopee/server        → internal/usecase/*
shopee/consumer      → internal/usecase/*, internal/infrastructure/kafka
shopee/scheduler     → (none, registers closures)
usecase/*            → internal/domain/* only
infrastructure/*     → internal/domain/*, internal (errors)
domain/*             → nothing (no imports from this module)
```

**Never allowed:**
- `internal/` importing from `sellchannel/`
- `usecase/` importing from `infrastructure/`
- `domain/` importing from anywhere in this module

## Error Handling

Two layers of errors:

**Generic** — `internal/errors.go`, package `internal`:
```go
internal.ErrNotFound
internal.ErrConflict
internal.ErrUnauthorized
internal.ErrInvalid
```

Use these in infrastructure when mapping low-level errors (e.g. `sql.ErrNoRows` → `internal.ErrNotFound`).

**Domain-specific** — `internal/domain/<name>/errors.go`:
```go
product.ErrNotActive
```

Use these in domain and usecase layers for business rule violations.

Usecases wrap errors with context:
```go
return fmt.Errorf("find product: %w", err)
```

Callers check sentinel errors with `errors.Is`:
```go
if errors.Is(err, internal.ErrNotFound) { ... }
```

## Port Interfaces

Each domain owns its own port interfaces in `port.go`. The domain defines what it needs — infrastructure implements it.

```go
// internal/domain/product/port.go
type Repository interface { ... }
type MarketplacePublisher interface { ... }
type EventPublisher interface { ... }
```

Infrastructure implements these interfaces implicitly (Go structural typing — no `implements` keyword needed).

## Composition Root

`sellchannel/shopee/shopee.go` is the only place where dependencies are wired together. It:
1. Creates infrastructure (DB, Redis connections)
2. Creates repositories (injects DB)
3. Creates use cases (injects repositories)
4. Creates server/consumer/scheduler (injects use cases)

No other file should do dependency injection. If you need to add a new dependency, add it here.

## Binaries

Each binary in `cmd/` does exactly three things:
1. Load config
2. Call `shopee.New(...)` to get the wired app
3. Start one component and block until signal

```go
// cmd/server/main.go
cfg := config.Load()
app, _ := shopee.New(cfg.Shopee, cfg.MySQL, cfg.Redis, cfg.Kafka)
app.Server().Run(":" + cfg.App.Port)
```

Build:
```bash
make build          # builds all three binaries into bin/
make run-server     # go run ./sellchannel/shopee/cmd/server
make run-consumer
make run-scheduler
```

Docker:
```bash
docker build -f sellchannel/shopee/cmd/server/Dockerfile -t shopee-server .
```

## Adding a New Sell Channel (e.g. Lazada)

1. Create `sellchannel/lazada/` mirroring the shopee structure
2. Implement `internal/domain/product.MarketplacePublisher` in `lazada/client/`
3. Add `Lazada lazada.Config` to `config/config.go`
4. Add Lazada env vars to `config.Load()`

The entire `internal/` layer (domain, usecase, infrastructure) is shared — do not duplicate it.

## Adding a New Domain (e.g. Promotion)

1. Create `internal/domain/promotion/entity.go` — define the struct
2. Create `internal/domain/promotion/port.go` — define Repository, MarketplacePublisher
3. Create `internal/domain/promotion/errors.go` — promotion-specific errors if any
4. Create `internal/infrastructure/mysql/promotion_repository.go` — implement the interface
5. Create `internal/usecase/promotion/sync.go` — business logic
6. Wire it in `sellchannel/shopee/shopee.go`

## Environment Variables

| Variable | Default | Description |
|---|---|---|
| `APP_ENV` | `development` | |
| `APP_PORT` | `8080` | |
| `SHOPEE_PARTNER_ID` | | |
| `SHOPEE_APP_KEY` | | |
| `SHOPEE_APP_SECRET` | | |
| `SHOPEE_BASE_URL` | `https://partner.shopeemobile.com` | |
| `MYSQL_DSN` | | e.g. `user:pass@tcp(host:3306)/db?parseTime=true` |
| `MYSQL_MAX_OPEN_CONNS` | `25` | |
| `MYSQL_MAX_IDLE_CONNS` | `5` | |
| `REDIS_ADDR` | `localhost:6379` | |
| `REDIS_PASSWORD` | | |
| `REDIS_DB` | `0` | |
| `KAFKA_BROKER` | `localhost:9092` | |
