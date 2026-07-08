# Shopee Channel Adapter

รับ order webhook จาก Shopee, enrich ผ่าน Shopee API แล้ว ship order อัตโนมัติผ่าน Kafka pipeline

---

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Project Structure](#project-structure)
- [Order Flow](#order-flow)
- [API Reference](#api-reference)
  - [GET /oauth/authorize](#get-oauthauthorize)
  - [GET /oauth/callback](#get-oauthcallback)
  - [POST /webhook](#post-webhook)
  - [GET /debug/order](#get-debugorder)
  - [GET /health](#get-health)
- [Database Schema](#database-schema)
- [Kafka Topics](#kafka-topics)
- [Idempotency Layers](#idempotency-layers)
- [Token Storage](#token-storage)
- [Makefile Reference](#makefile-reference)

---

## Architecture Overview

```
Shopee Platform
      │
      │  Signed order webhook (POST /webhook)
      ▼
┌─────────────────────────────────────────────────────┐
│  HTTP Server  (cmd/server, port 8085)               │
│  ┌──────────────────────────────────────────────┐   │
│  │  shopeeWebhookAuth middleware                │   │
│  │  HMAC-SHA256(appSecret, partnerID+path+body) │   │
│  └──────────────────────────────────────────────┘   │
│                │                                    │
│         Verify signature                            │
│         Token check (Redis → Postgres)              │
│         Redis dedup SetNX                           │
│                │                                    │
│         Publish → shopee.order.raw                 │
└─────────────────────────────────────────────────────┘
                 │
    ┌────────────▼──────────┐     ┌──────────────────────┐
    │  Kafka Broker         │     │  Redis               │
    │  shopee.order.raw     │     │  - dedup cache       │
    │  shopee.order.detail  │     │  - token cache       │
    └────────────┬──────────┘     │  - poll seen cache   │
                 │                └──────────────────────┘
      ┌──────────▼──────────────────────────────┐
      │  Order Consumer  (cmd/consumer)          │
      │  ├── GetOrderDetail (Shopee API)         │
      │  ├── Publish → shopee.order.detail       │
      │  └── Enqueue fulfillment (READY_TO_SHIP) │
      └──────────────────────────┬──────────────┘
                                 │
                    ┌────────────▼────────────┐
                    │  PostgreSQL             │
                    │  shopee_tokens          │
                    │  shopee_fulfillment     │
                    └────────────┬────────────┘
                                 │
      ┌──────────────────────────▼──────────────────────┐
      │  Fulfillment Scheduler  (cmd/fulfillment-scheduler) │
      │  @every 5m                                      │
      │  ├── ListPending from shopee_fulfillment        │
      │  ├── GetShippingParameter (Shopee API)          │
      │  ├── ShipOrder (Shopee API)                     │
      │  ├── GetTrackingNumber (Shopee API)             │
      │  └── MarkShipped in DB                          │
      └─────────────────────────────────────────────────┘

 Fallback Polling (cmd/scheduler — @every 5m, window=10m):
      ├── FindAllShopIDs from shopee_tokens
      ├── GetOrderList (Shopee API) per shop
      ├── Redis SetNX dedup (shopee:poll:seen:{order_sn})
      └── Publish missed orders → shopee.order.raw
```

**Idempotency Layers:**

| Layer | Mechanism | Key | TTL |
|---|---|---|---|
| Webhook dedup | Redis SetNX | `shopee:webhook:dedup:{order_sn}:{timestamp}` | 24h |
| Poll dedup | Redis SetNX | `shopee:poll:seen:{order_sn}` | 10m |
| Token check | tokenstore.Get | `shopee:{shop_id}:access_token` | token lifetime |
| Fulfillment queue | Postgres `ON CONFLICT DO NOTHING` | `(shop_id, order_sn)` | permanent |

---

## Tech Stack

| Category | Library | Version |
|---|---|---|
| HTTP Framework | [gin-gonic/gin](https://github.com/gin-gonic/gin) | v1.12 |
| Kafka Client | [twmb/franz-go](https://github.com/twmb/franz-go) | v1.21 |
| Database | PostgreSQL | 16 |
| DB Driver | [jackc/pgx](https://github.com/jackc/pgx) | v5 |
| Cache | Redis | 7 |
| Redis Client | [redis/go-redis](https://github.com/redis/go-redis) | v9 |
| Scheduler | [robfig/cron](https://github.com/robfig/cron) | v3 |
| HTTP Client | [go-resty/resty](https://github.com/go-resty/resty) | v2 |
| JSON | [bytedance/sonic](https://github.com/bytedance/sonic) | v1.15 |
| Logger | go.uber.org/zap | v1.28 |

---

## Prerequisites

| Tool | Version | Purpose |
|---|---|---|
| Go | 1.26+ | Runtime |
| Docker + Compose | latest | Local infra (Kafka, Postgres, Redis) |
| make | — | Build automation |
| ngrok (optional) | latest | Expose local server ให้ Shopee webhook เข้าถึงได้ |

---

## Quick Start

```bash
# 1. Clone
git clone <repo-url>
cd phoenix-multichannel-mkp

# 2. Copy และแก้ไข config
cp .env.example .env   # ใส่ SHOPEE_PARTNER_ID, SHOPEE_APP_KEY, SHOPEE_APP_SECRET

# 3. Start infra (Kafka, Postgres, Redis)
docker compose -f deployments/docker-compose.yml up -d

# 4. ทำ OAuth authorize เพื่อ get token ร้านค้า
open http://localhost:8085/oauth/authorize
# → redirect ไป Shopee login → callback จะ save token ใน Postgres + Redis

# 5. Start all services (แต่ละ terminal)
make run-server                    # HTTP server (port 8085)
make run-consumer                  # Kafka consumer
make run-scheduler                 # Order poll scheduler
make run-fulfillment-scheduler     # Fulfillment scheduler
```

ตรวจสอบ:

```bash
curl http://localhost:8085/health

# Redpanda Console
open http://localhost:8081
```

---

## Configuration

`.env` (ไม่ถูก commit)

```env
# App
APP_ENV=development
APP_PORT=8085

# Shopee Open Platform
SHOPEE_PARTNER_ID=1236084
SHOPEE_APP_KEY=shpk...
SHOPEE_APP_SECRET=shpk...          # ต้องเท่ากับ APP_KEY (Shopee ใช้ key เดียว)
SHOPEE_BASE_URL=https://openplatform.sandbox.test-stable.shopee.sg
SHOPEE_OAUTH_REDIRECT_URL=https://<your-ngrok>.ngrok-free.dev/oauth/callback
SHOPEE_WEBHOOK_VERIFY=false        # true ใน production

# Order Poll Scheduler
SHOPEE_POLL_SPEC="@every 5m"
SHOPEE_POLL_WINDOW=10m

# Fulfillment Scheduler
SHOPEE_FULFILLMENT_SPEC="@every 5m"
SHOPEE_FULFILLMENT_LIMIT=20

# Infra
POSTGRES_DSN=postgres://postgres:password@localhost:5432/marketplace?sslmode=disable
REDIS_ADDR=localhost:6379
KAFKA_BROKER=localhost:9092
KAFKA_GROUP_ID=marketplace-sync
```

| Variable | Default | หมายเหตุ |
|---|---|---|
| `SHOPEE_PARTNER_ID` | — | จาก Shopee Partner Portal |
| `SHOPEE_APP_KEY` | — | Live / Test API Partner Key |
| `SHOPEE_APP_SECRET` | — | ต้องเท่ากับ `APP_KEY` (Shopee ใช้ key เดียวกัน) |
| `SHOPEE_BASE_URL` | `https://partner.shopeemobile.com` | ใช้ sandbox URL ตอน dev |
| `SHOPEE_WEBHOOK_VERIFY` | `false` | ตั้งเป็น `true` ใน production เสมอ |
| `SHOPEE_POLL_WINDOW` | `10m` | ควรมากกว่า poll interval เพื่อ overlap |
| `SHOPEE_FULFILLMENT_LIMIT` | `20` | max orders ต่อ scheduler run |

**Sandbox vs Production:**

| | Sandbox | Production |
|---|---|---|
| `SHOPEE_BASE_URL` | `https://openplatform.sandbox.test-stable.shopee.sg` | `https://partner.shopeemobile.com` |
| `SHOPEE_WEBHOOK_VERIFY` | `false` | `true` |
| `SHOPEE_PARTNER_ID` | sandbox partner id | production partner id |

---

## Project Structure

```
sellchannel/shopee/
│
├── cmd/
│   ├── server/                  # HTTP server binary (webhook + OAuth)
│   ├── consumer/                # Kafka consumer binary (order enrichment)
│   ├── scheduler/               # Order poll scheduler binary (fallback)
│   └── fulfillment-scheduler/   # Fulfillment scheduler binary (ship orders)
│
├── client/                      # Shopee API client
│   ├── auth.go                  # OAuth: AuthURL(), GetAccessToken(), RefreshToken()
│   ├── order.go                 # GetOrderList(), GetOrderDetail()
│   └── fulfillment.go           # GetShippingParameter(), ShipOrder*(), GetTrackingNumber()
│
├── server/                      # Gin HTTP handlers
│   ├── router.go                # Server struct + route registration
│   ├── webhook.go               # POST /webhook — dedup → token check → Kafka publish
│   ├── webhook_verify.go        # HMAC-SHA256 signature verification middleware
│   ├── webhook_dto.go           # OrderWebhookRequest, OrderRawEvent structs
│   ├── oauth.go                 # GET /oauth/authorize, GET /oauth/callback
│   ├── oauth_dto.go             # OAuthCallbackRequest struct
│   ├── debug_order.go           # GET /debug/order (simulate webhook, non-production)
│   ├── health.go                # GET /health
│   └── middleware.go            # Request logger
│
├── consumer/
│   ├── order.go                 # OrderConsumer: raw → GetOrderDetail → detail + fulfillment
│   └── product.go               # ProductConsumer (stub)
│
├── scheduler/
│   └── order_poll.go            # OrderPollJob: poll GetOrderList → publish missed orders
│
├── fulfillmentscheduler/
│   └── scheduler.go             # FulfillmentScheduler: cron wrapper
│
├── fulfillmentsync/
│   └── syncer.go                # Syncer: shopee_logistics + own_fleet ship logic
│
├── tokenstore/
│   └── store.go                 # Token R/W: Postgres (source of truth) + Redis (cache)
│
├── config/                      # Shopee-specific config (embedded in shopee.Config)
├── config.go                    # shopee.Config struct
└── shopee.go                    # Composition root — wires all dependencies
```

```
internal/
├── domain/order/                # Order entity + port interfaces
├── usecase/order/               # ProcessWebhookUsecase
└── infrastructure/
    ├── postgres/
    │   ├── shopee_token_repository.go       # shopee_tokens table
    │   ├── shopee_fulfillment_repository.go # shopee_fulfillment table
    │   └── order_repository.go              # orders table (schema pending)
    ├── redis/                   # Redis client + SetNX / GetString helpers
    └── kafka/                   # Producer + Consumer wrappers (franz-go)
```

---

## Order Flow

### 1. Webhook → Kafka

```
Shopee Platform
    │  POST /webhook
    │  Authorization: <HMAC-SHA256 hex>
    ▼
shopeeWebhookAuth middleware
    ├── read raw body (buffer for downstream)
    ├── SHOPEE_WEBHOOK_VERIFY=true  → verify HMAC-SHA256(appSecret, partnerID+path+body)
    └── SHOPEE_WEBHOOK_VERIFY=false → skip (sandbox/dev)
    │
    ▼
handleOrderWebhook()
    ├── Bind JSON → OrderWebhookRequest
    ├── return 200 OK immediately  ← Shopee requires 2xx within timeout
    └── go func() {
            ├── Redis SetNX shopee:webhook:dedup:{order_sn}:{timestamp} EX 24h
            │     └── !isNew → duplicate, return
            ├── tokenstore.Get(shopID)
            │     └── err → shop not authorized, return (ไม่ publish ไป Kafka)
            └── Kafka Publish → shopee.order.raw
                  key   = order_sn
                  value = { shop_id, order_sn, status, timestamp }
        }
```

### 2. Order Consumer → Kafka

```
shopee.order.raw
    │
    ▼
OrderConsumer.Start()
    └── process(event)
          ├── tokenstore.Get(shop_id)  → get access token
          ├── shopeeClient.GetOrderDetail(shop_id, token, [order_sn])
          ├── Kafka Publish → shopee.order.detail
          │     key   = order_sn
          │     value = full Shopee order object
          └── order_status == "READY_TO_SHIP"
                └── fulfillmentRepo.Enqueue(shop_id, order_sn, "shopee_logistics", "")
                      INSERT INTO shopee_fulfillment ON CONFLICT DO NOTHING
```

### 3. Fallback Poll → Kafka (ถ้า webhook พลาด)

```
OrderPollJob.Run()  (@every 5m, window=10m)
    ├── tokenRepo.FindAllShopIDs()
    └── per shop:
          ├── tokens.Get(shop_id)
          ├── shopeeClient.GetOrderList(shop_id, token, now-10m, now)
          └── per order:
                ├── Redis SetNX shopee:poll:seen:{order_sn} EX 10m
                │     └── !isNew → already published by webhook, skip
                └── Kafka Publish → shopee.order.raw  ← same topic, same consumer
```

### 4. Fulfillment Scheduler → Shopee API

```
FulfillmentScheduler  (@every 5m, limit=20)
    └── syncer.SyncPending(ctx, 20)
          ├── fulfillmentRepo.ListPending(limit=20)  ← WHERE sync_status='pending'
          └── per row:
                ├── tokens.Get(shop_id)
                │
                ├── delivery_type = "shopee_logistics"
                │     ├── GetShippingParameter(order_sn)  → pickup address + timeslots
                │     ├── ShipOrderPickup(order_sn, address_id, pickup_time_id)
                │     ├── GetTrackingNumber(order_sn)
                │     └── fulfillmentRepo.MarkShipped(id, tracking_number)
                │
                └── delivery_type = "own_fleet"
                      ├── ShipOrderNonIntegrated(order_sn, tracking_number)
                      └── fulfillmentRepo.MarkShipped(id, tracking_number)
```

---

## API Reference

### GET /oauth/authorize

Redirect merchant ไป Shopee authorization page

```bash
open http://localhost:8085/oauth/authorize
```

→ redirect ไป `{SHOPEE_BASE_URL}/api/v2/shop/auth_partner?...` พร้อม HMAC-signed URL

### GET /oauth/callback

Shopee redirect กลับมาพร้อม authorization code — server แลก code → token แล้ว save

```
GET /oauth/callback?code=<auth_code>&shop_id=<shop_id>
```

**Response 200:**

```json
{
  "shop_id": 225997847,
  "expires_in": 14400,
  "message": "token stored successfully"
}
```

| Code | เงื่อนไข |
|---|---|
| `200 OK` | token stored ใน Postgres + Redis |
| `400 Bad Request` | missing code หรือ shop_id |
| `401 Unauthorized` | Shopee ตอบ error (invalid_code, invalid_partner_id ฯลฯ) |
| `502 Bad Gateway` | Shopee API error |

> **Authorization code ใช้ได้ครั้งเดียวและหมดอายุใน 10 นาที** ห้ามใช้ code เดิมซ้ำหรือ copy ไป curl เพราะ server exchange ให้อัตโนมัติแล้ว

### POST /webhook

รับ Shopee push notification — return 200 OK ทันที, process async

```
POST /webhook
Authorization: <HMAC-SHA256 signature from Shopee>
Content-Type: application/json
```

```json
{
  "code": 3,
  "timestamp": 1783421265,
  "shop_id": 225997847,
  "data": {
    "ordersn": "2607073G09MYPG",
    "status": "READY_TO_SHIP"
  }
}
```

**Response 200 (always):**

```json
{ "status": "ok" }
```

> Shopee ต้องการ 2xx เสมอ ไม่เช่นนั้น Shopee จะ retry ซ้ำ — handler return 200 ทันทีแล้ว process ใน goroutine

**Push codes ที่รองรับ:**

| Code | Event |
|---|---|
| `3` | Order status update |
| `4` | Product status update |

### GET /debug/order

Simulate webhook โดย fetch order จาก Shopee API แล้ว publish ไป Kafka (ใช้ได้เฉพาะ dev)

```bash
curl "http://localhost:8085/debug/order?shop_id=225997847&order_sn=2607073G09MYPG"
```

**Response 200:**

```json
{
  "shop_id": 225997847,
  "order_sn": "2607073G09MYPG",
  "kafka_topic": "shopee.order.raw",
  "kafka_published": true,
  "message": "raw event published — consumer will fetch order detail from Shopee API"
}
```

| Code | เงื่อนไข |
|---|---|
| `200 OK` | published to Kafka |
| `400 Bad Request` | missing shop_id หรือ order_sn |
| `401 Unauthorized` | shop ไม่มี token (ต้อง OAuth ก่อน) |
| `500 Internal Server Error` | Kafka error |

### GET /health

```bash
curl http://localhost:8085/health
# → 200 OK  { "status": "ok" }
```

---

## Database Schema

### `shopee_tokens`

OAuth token ต่อ shop — Postgres เป็น source of truth, Redis เป็น fast-path cache

```sql
CREATE TABLE shopee_tokens (
    shop_id       BIGINT      PRIMARY KEY,
    access_token  TEXT        NOT NULL,
    refresh_token TEXT        NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    expired_at    TIMESTAMPTZ NOT NULL
);
```

| Column | หมายเหตุ |
|---|---|
| `shop_id` | Shopee Shop ID (PK) |
| `access_token` | ใช้ sign API requests |
| `refresh_token` | ใช้ renew access token |
| `expired_at` | Shopee token อายุ ~4 ชั่วโมง — ต้อง re-authorize ก่อนหมดอายุ |

### `shopee_fulfillment`

Queue ของ order ที่รอ ship

```sql
CREATE TABLE shopee_fulfillment (
    id              BIGSERIAL   PRIMARY KEY,
    shop_id         BIGINT      NOT NULL,
    order_sn        TEXT        NOT NULL,
    delivery_type   TEXT        NOT NULL DEFAULT 'shopee_logistics',
    tracking_number TEXT        NOT NULL DEFAULT '',
    sync_status     TEXT        NOT NULL DEFAULT 'pending',
    last_error      TEXT        NOT NULL DEFAULT '',
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (shop_id, order_sn)
);
```

| Column | Values | หมายเหตุ |
|---|---|---|
| `delivery_type` | `shopee_logistics` / `own_fleet` | shopee_logistics = Shopee จัด pickup ให้ |
| `sync_status` | `pending` / `shipped` | scheduler retry rows ที่ยัง `pending` |
| `tracking_number` | string | shopee_logistics = Shopee return หลัง ship, own_fleet = seller provide |
| `last_error` | string | error ล่าสุด (row ยังเป็น pending, scheduler จะ retry) |

**ดู fulfillment queue:**

```sql
-- pending orders
SELECT shop_id, order_sn, delivery_type, last_error, updated_at
FROM shopee_fulfillment
WHERE sync_status = 'pending'
ORDER BY updated_at;

-- shipped orders
SELECT shop_id, order_sn, tracking_number, updated_at
FROM shopee_fulfillment
WHERE sync_status = 'shipped'
ORDER BY updated_at DESC;
```

**Schema สร้างอัตโนมัติตอน startup** ผ่าน `EnsureSchema()` — ไม่ต้องรัน migration มือ

---

## Kafka Topics

| Topic | Key | Producer | Consumer | หมายเหตุ |
|---|---|---|---|---|
| `shopee.order.raw` | `order_sn` | webhook handler, order poll | OrderConsumer | lightweight event — แค่ shop_id + order_sn + status |
| `shopee.order.detail` | `order_sn` | OrderConsumer | downstream (TBD) | full order detail จาก Shopee GetOrderDetail API |

**ดู messages ใน Redpanda Console:** `http://localhost:8081`

**shopee.order.raw payload:**

```json
{
  "shop_id": 225997847,
  "order_sn": "2607073G09MYPG",
  "status": "READY_TO_SHIP",
  "timestamp": 1783421265
}
```

---

## Idempotency Layers

### Webhook dedup

```
Redis SetNX  shopee:webhook:dedup:{order_sn}:{timestamp}  EX 24h
```

Shopee retry webhook ด้วย `timestamp` เดิมถ้าไม่ได้รับ 2xx → key เดิม → skip

### Poll dedup

```
Redis SetNX  shopee:poll:seen:{order_sn}  EX 10m
```

ป้องกัน order จาก webhook path ถูก publish ซ้ำโดย poll fallback

### Fulfillment dedup

```sql
INSERT INTO shopee_fulfillment (shop_id, order_sn, ...)
ON CONFLICT (shop_id, order_sn) DO NOTHING
```

Consumer call Enqueue ได้กี่ครั้งก็ได้ — row สร้างแค่ครั้งเดียว

---

## Token Storage

`tokenstore.Store` ใช้ two-layer storage:

```
tokenstore.Get(ctx, shopID)
    ├── Redis GET shopee:{shop_id}:access_token   ← fast path (token lifetime TTL)
    │     └── HIT → return immediately
    └── MISS → Postgres SELECT shopee_tokens WHERE shop_id = ?
                  ├── token expired    → log WARNING, return token (caller handles)
                  ├── expiring < 24h   → log WARNING (ควร renew)
                  └── re-warm Redis cache → return token
```

**Token expiry:** Shopee access token อายุ ~4 ชั่วโมง — ต้อง re-authorize ผ่าน `/oauth/authorize` ก่อนหมดอายุ

---

## Makefile Reference

```bash
# Run (Shopee)
make run-server                     # HTTP server (port 8085)
make run-consumer                   # Kafka order consumer
make run-scheduler                  # Order poll scheduler (@every 5m)
make run-fulfillment-scheduler      # Fulfillment scheduler (@every 5m)

# Build
make build                          # Build all binaries to bin/

# Docker images
make docker-build                   # Build shopee-server, shopee-consumer, shopee-scheduler

# Quality
make test                           # Run all tests
make test-cover                     # Tests + HTML coverage report
make lint                           # golangci-lint
make vet                            # go vet
make fmt                            # go fmt
make tidy                           # go mod tidy + verify
make check                          # fmt + vet + test

# Cleanup
make clean                          # Remove bin/ coverage.out coverage.html
```

---

## Order Status Reference

```
UNPAID → READY_TO_SHIP → PROCESSED → SHIPPED → COMPLETED
                                    ↘ CANCELLED / IN_CANCEL
```

| Status | ความหมาย | Action ใน system |
|---|---|---|
| `UNPAID` | รอชำระ | รับ webhook, ไม่ enqueue fulfillment |
| `READY_TO_SHIP` | ชำระแล้ว รอ ship | OrderConsumer enqueue fulfillment |
| `PROCESSED` | ship_order สำเร็จ | fulfillment MarkShipped |
| `SHIPPED` | courier รับพัสดุแล้ว | — |
| `COMPLETED` | ผู้ซื้อได้รับสินค้า | — |
| `CANCELLED` | ยกเลิก | — |
