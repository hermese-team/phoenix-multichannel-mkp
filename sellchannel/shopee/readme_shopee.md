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
- [Safety-Net Scan](#safety-net-scan)
- [Classifier](#classifier)
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
      │  Signed order webhook (POST /webhook)      Fallback: Order Poll (*/5 * * * *)
      ▼                                                         │
┌─────────────────────────────────────────────────────┐        │
│  HTTP Server  (cmd/server, port 8085)               │        │
│  ┌──────────────────────────────────────────────┐   │        │
│  │  shopeeWebhookAuth middleware                │   │        │
│  │  HMAC-SHA256(appSecret, partnerID+path+body) │   │        │
│  └──────────────────────────────────────────────┘   │        │
└─────────────────────┬───────────────────────────────┘        │
                      │                                         │
             ┌────────▼─────────────────────────────────────────┘
             │
             ▼
┌─────────────────────────────────────────────────────┐
│  intake.Accept()  (W — Webhook Intake component)    │
│  ├── Redis SetNX  shopee:[webhook|poll]:dedup:...   │
│  ├── Redis SADD   safety-net:shopee:processed       │
│  └── Kafka Publish → raw.accepted.shopee.v1.dev     │
└─────────────────────┬───────────────────────────────┘
                      │
        ┌─────────────▼──────────────────┐
        │  raw.accepted.shopee.v1.dev    │
        └─────────────┬──────────────────┘
                      │
       ┌──────────────▼──────────────────────────────────┐
       │  Classifier Consumer  (cmd/consumer)             │
       │  Classify(pushCode, status) → canonical EventType│
       │  ORDER_CREATED / ORDER_READY_TO_SHIP / ...       │
       └──────────────┬──────────────────────────────────┘
                      │
        ┌─────────────▼────────────────────┐
        │  order.ingest.shopee.v1.dev      │
        └─────────────┬────────────────────┘
                      │
       ┌──────────────▼──────────────────────────────────┐
       │  Order Consumer  (cmd/consumer)                  │
       │  ├── Coalesce 2s window (deduplicate bursts)     │
       │  ├── Bulk GetOrderDetail (Shopee API, ≤50/call)  │
       │  ├── Enqueue fulfillment (READY_TO_SHIP)         │
       │  └── Publish → order.enriched.v1.dev             │
       │  (failures → order.dlq.v1.dev)                   │
       └──────────────┬──────────────────────────────────┘
                      │
        ┌─────────────▼──────────────────┐
        │  order.enriched.v1.dev         │
        └─────────────┬──────────────────┘
                      │
       ┌──────────────▼──────────────────────────────────┐
       │  Ingestion Consumer  (cmd/consumer)              │
       │  ├── Normalize + PostgreSQL idempotency check    │
       │  ├── Upsert → orders (monthly partitioned)       │
       │  └── Publish → order.received.v1.dev             │
       └──────────────┬──────────────────────────────────┘
                      │
        ┌─────────────▼──────────────────┐
        │  order.received.v1.dev         │  ← downstream
        └────────────────────────────────┘

 order.lifecycle.v1.dev  ← history/audit-log topic (published by separate consumer)
 order.retry.v1.dev      ← retry topic

 Safety-Net Scan (cmd/scheduler — 2,17,32,47 * * * *):
      ├── GetOrderList per shop (last 30m)
      ├── SMISMEMBER safety-net:shopee:processed
      ├── Re-inject missed orders → raw.accepted.shopee.v1.dev
      └── Publish governance event → safety-net.scan.complete.v1

 Fulfillment Scheduler (cmd/fulfillment-scheduler — @every 5m):
      ├── ListPending from shopee_fulfillment
      ├── GetShippingParameter / ShipOrder / GetTrackingNumber (Shopee API)
      └── MarkShipped in DB
```

**Redis keys:**

| Key | Type | TTL | Purpose |
|---|---|---|---|
| `shopee:webhook:dedup:{order_sn}:{timestamp}` | String | 24h | Webhook dedup |
| `shopee:poll:seen:{order_sn}` | String | 24h | Poll dedup |
| `safety-net:shopee:processed` | Set | 30m | Safety-net membership |
| `shopee:{shop_id}:access_token` | String | token lifetime | Token cache |
| `shopee:poll:lease` | String | 4m | Distributed poll lock |
| `shopee:scan:lease` | String | 14m | Distributed scan lock |

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
make run-consumer                  # Kafka consumers (classifier + enricher + ingestion)
make run-scheduler                 # Order poll + safety-net scan
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
# Poll runs at :00,:05,:10,...,:55 (fixed-minute boundaries)
SHOPEE_POLL_SPEC="*/5 * * * *"
SHOPEE_POLL_WINDOW=10m

# Safety-Net Scan Scheduler
# Scan runs at :02,:17,:32,:47 — always 2 min after a poll tick so
# intake.Accept SADD has time to persist before SMISMEMBER cross-reference runs
SHOPEE_SCAN_SPEC="2,17,32,47 * * * *"

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
| `SHOPEE_POLL_SPEC` | `*/5 * * * *` | fixed-minute เพื่อให้ scan stagger ได้แม่นยำ |
| `SHOPEE_SCAN_SPEC` | `2,17,32,47 * * * *` | รัน 2 นาทีหลัง poll เสมอ |
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
│   ├── consumer/                # Kafka consumer binary (classifier + enricher + ingestion)
│   ├── scheduler/               # Order poll + safety-net scan binary
│   └── fulfillment-scheduler/   # Fulfillment scheduler binary (ship orders)
│
├── client/                      # Shopee API client
│   ├── auth.go                  # OAuth: AuthURL(), GetAccessToken(), RefreshToken()
│   ├── order.go                 # GetOrderList() (incl. response_optional_fields=order_status)
│   │                            # GetOrderDetail()
│   └── fulfillment.go           # GetShippingParameter(), ShipOrder*(), GetTrackingNumber()
│
├── intake/                      # W (Webhook Intake) component — A→W interface
│   └── intake.go                # Intake.Accept(): dedup SetNX + safety-net SADD + Kafka publish
│                                # RawEvent struct, OrderRawTopic, SafetyNetKey, DedupTTL constants
│
├── server/                      # Gin HTTP handlers (A — Channel Adapter)
│   ├── router.go                # Server struct + route registration
│   ├── webhook.go               # POST /webhook — auth → intake.Accept()
│   ├── webhook_verify.go        # HMAC-SHA256 signature verification middleware
│   ├── webhook_dto.go           # OrderWebhookRequest struct
│   ├── oauth.go                 # GET /oauth/authorize, GET /oauth/callback
│   ├── oauth_dto.go             # OAuthCallbackRequest struct
│   ├── debug_order.go           # GET /debug/order — simulate push via intake.Accept()
│   ├── health.go                # GET /health
│   └── middleware.go            # Request logger
│
├── classifier/                  # Step 4: canonical event type mapping
│   ├── consumer.go              # Classifier: order.raw.accepted.v1 → order.ingest.shopee.v1
│   └── classifier.go            # Classify(pushCode int, status string) → EventType string
│
├── consumer/
│   ├── order.go                 # OrderConsumer: ingest → GetOrderDetail → enriched + fulfillment
│   ├── ingestion.go             # IngestionConsumer: enriched → orders DB upsert → lifecycle event
│   └── product.go               # ProductConsumer (stub)
│
├── scheduler/
│   ├── scheduler.go             # Scheduler: cron wrapper for poll + scan jobs
│   ├── order_poll.go            # OrderPollJob: GetOrderList → SMISMEMBER → intake.Accept()
│   └── safety_net_scan.go       # SafetyNetScanJob: audit pipeline, re-inject missed orders
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
    │   ├── shopee_poll_cursor_repository.go # shopee_poll_cursor table
    │   ├── safety_net_scan_repository.go    # safety_net_scan_results table
    │   └── order_repository.go              # orders table (monthly partitioned)
    ├── redis/                   # Redis client + SetNX / SADD / SMIsMember helpers
    └── kafka/                   # Producer + Consumer wrappers (franz-go)
```

---

## Order Flow

### 1. Webhook → Intake → Kafka

```
Shopee Platform
    │  POST /webhook
    │  Authorization: <HMAC-SHA256 hex>
    ▼
shopeeWebhookAuth middleware
    ├── SHOPEE_WEBHOOK_VERIFY=true  → verify HMAC-SHA256(appSecret, partnerID+path+body)
    └── SHOPEE_WEBHOOK_VERIFY=false → skip (sandbox/dev)
    │
    ▼
handleOrderWebhook()
    ├── Bind JSON → OrderWebhookRequest
    ├── return 200 OK immediately  ← Shopee requires 2xx within timeout
    └── go func() {
            intake.Accept(ctx, "shopee:webhook:dedup:{sn}:{ts}", RawEvent{...})
                ├── Redis SetNX shopee:webhook:dedup:{order_sn}:{timestamp} EX 24h
                │     └── !isNew → duplicate → return
                ├── Redis SADD safety-net:shopee:processed {order_sn} (TTL 30m)
                └── Kafka Publish → order.raw.accepted.v1
                      key   = order_sn
                      value = { shop_id, order_sn, status, code, timestamp }
        }
```

### 2. Fallback Poll → Intake → Kafka (ถ้า webhook พลาด)

```
OrderPollJob.Run()  (*/5 * * * *  window=30m)
    ├── tokenRepo.FindAllShopIDs()
    └── per shop:
          ├── tokens.Get(shop_id)
          ├── shopeeClient.GetOrderList(shop_id, token, now-30m, now)
          │     (incl. response_optional_fields=order_status)
          │
          ├── Layer 1: SMISMEMBER safety-net:shopee:processed [all order_sn]
          │     └── in set → already handled by webhook → skip
          │
          └── Layer 2: per remaining order:
                intake.Accept(ctx, "shopee:poll:seen:{sn}", RawEvent{PushCode:0, Status:...})
                    ├── Redis SetNX shopee:poll:seen:{order_sn} EX 24h → skip if seen
                    ├── Redis SADD safety-net:shopee:processed {order_sn}
                    └── Kafka Publish → order.raw.accepted.v1
```

**Cron stagger:** Poll (`*/5 * * * *`) runs at fixed-minute boundaries (:00, :05, …). Scan (`2,17,32,47 * * * *`) runs 2 minutes later, guaranteeing SADD is persisted before SMISMEMBER. Using `@every` for both would cause a race since both timers start at scheduler init time.

### 3. Classifier → Kafka

```
order.raw.accepted.v1
    │
    ▼
Classifier Consumer  (group: shopee-classifier)
    └── classify(raw)
          ├── Classify(raw.PushCode, raw.Status) → canonical EventType
          └── Kafka Publish → order.ingest.shopee.v1
                { channel, event_type, shop_id, order_sn, status, code, timestamp }
```

### 4. Order Enricher → Kafka

```
order.ingest.shopee.v1
    │
    ▼
OrderConsumer  (group: shopee-order-enricher)
    └── process(event)
          ├── Coalesce 2s window — merge duplicate order_sn pushes
          ├── Bulk GetOrderDetail(shop_id, token, [order_sn, ...])  ← ≤50 per call, 5 RPS
          ├── On success → Kafka Publish → order.enriched.v1
          ├── On permanent failure (3 retries) → Kafka Publish → order.enriched.shopee.dlq.v1
          └── event_type == ORDER_READY_TO_SHIP
                └── fulfillmentRepo.Enqueue(shop_id, order_sn, "shopee_logistics", "")
                      INSERT INTO shopee_fulfillment ON CONFLICT DO NOTHING
```

### 5. Ingestion Consumer → Database

```
order.enriched.v1
    │
    ▼
IngestionConsumer  (group: shopee-order-ingestion)
    └── upsert(enriched)
          ├── orderRepo.Upsert(ctx, order)   ← find-or-create across partitions
          └── Kafka Publish → order.lifecycle.v1
```

### 6. Fulfillment Scheduler → Shopee API

```
FulfillmentScheduler  (@every 5m, limit=20)
    └── syncer.SyncPending(ctx, 20)
          ├── fulfillmentRepo.ListPending(limit=20)  ← WHERE sync_status='pending'
          └── per row:
                ├── delivery_type = "shopee_logistics"
                │     ├── GetShippingParameter(order_sn)
                │     ├── ShipOrderPickup(order_sn, address_id, pickup_time_id)
                │     ├── GetTrackingNumber(order_sn)
                │     └── fulfillmentRepo.MarkShipped(id, tracking_number)
                └── delivery_type = "own_fleet"
                      ├── ShipOrderNonIntegrated(order_sn, tracking_number)
                      └── fulfillmentRepo.MarkShipped(id, tracking_number)
```

---

## Safety-Net Scan

`SafetyNetScanJob` รัน `2,17,32,47 * * * *` (2 นาทีหลัง poll ทุก tick) เพื่อตรวจว่า webhook pipeline พลาด order ไหนในช่วง 30 นาทีล่าสุด

```
SafetyNetScanJob.Run()
    ├── Acquire distributed lease (Redis SetNX shopee:scan:lease EX 14m)
    └── per shop:
          ├── GetOrderList(shop_id, token, now-30m, now)
          ├── SMISMEMBER safety-net:shopee:processed [all order_sn]
          │     ├── in set  → matched → matchedCount++
          │     └── not in set → missed → missedCount++
          │           └── Publish → order.raw.accepted.v1  (re-inject)
          ├── Kafka Publish → safety-net.scan.complete.v1
          │     { channel, shop_id, window_from, window_to,
          │       total, matched, missed, reinjected, duration_ms }
          └── Persist → safety_net_scan_results table
```

**ตัวชี้วัดสุขภาพ:** `missed=0` ทุก scan = webhook pipeline ปกติ. `missed>0` = scan re-inject ให้แล้ว — order จะถูก process ผ่าน pipeline ปกติ.

**Cron stagger design:**
- Poll รันที่ `:00,:05,:10,...` → SADD `safety-net:shopee:processed`
- Scan รันที่ `:02,:17,:32,:47` → SMISMEMBER อ่าน Redis หลัง poll เสร็จ
- `@every 5m` + `@every 15m` จะ race: ทั้งคู่ start ที่ `scheduler init time` อาจรันพร้อมกัน

---

## Classifier

`classifier.Consumer` อ่านจาก `order.raw.accepted.v1` และ map Shopee push code + order status ไปเป็น canonical `EventType` ก่อน publish ไป `order.ingest.shopee.v1`

เหตุผล: downstream consumers ไม่ควรรู้จัก Shopee push code โดยตรง — ใช้ canonical type เพื่อ decouple

**Canonical EventTypes:**

| PushCode | Status | EventType |
|---|---|---|
| `3` | `UNPAID` | `ORDER_CREATED` |
| `3` | `READY_TO_SHIP` | `ORDER_READY_TO_SHIP` |
| `3` | `PROCESSED` | `ORDER_PROCESSED` |
| `3` | `SHIPPED` | `ORDER_SHIPPED` |
| `3` | `COMPLETED` | `ORDER_COMPLETED` |
| `3` | `CANCELLED` | `ORDER_CANCELLED` |
| `3` | `IN_CANCEL` | `ORDER_IN_CANCEL` |
| `0` (poll) | any | same status-based mapping |
| other | any | `ORDER_STATUS_UPDATED` (fallback) |

**GetOrderList `order_status` field:** Shopee ต้องการ `response_optional_fields=order_status` เพื่อ return `OrderStatus` — ถ้าไม่ใส่ status จะเป็น `""` และ classifier จะได้ `ORDER_STATUS_UPDATED` เสมอ. `client.GetOrderList` ใส่ param นี้ให้แล้ว.

**IngestEvent payload** (published to `order.ingest.shopee.v1`):

```json
{
  "channel": "shopee",
  "event_type": "ORDER_READY_TO_SHIP",
  "shop_id": 225997847,
  "order_sn": "26070985HRXASU",
  "status": "READY_TO_SHIP",
  "code": 3,
  "timestamp": 1783421265
}
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

รับ Shopee push notification — return 200 OK ทันที, process ผ่าน `intake.Accept()` async

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

Simulate webhook push โดย publish `RawEvent` ไป `order.raw.accepted.v1` ผ่าน `intake.Accept()` (ใช้ได้เฉพาะ dev)

```bash
curl "http://localhost:8085/debug/order?shop_id=225997847&order_sn=2607073G09MYPG"
```

**Response 200:**

```json
{
  "shop_id": 225997847,
  "order_sn": "2607073G09MYPG",
  "kafka_topic": "order.raw.accepted.v1",
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
SELECT shop_id, order_sn, delivery_type, last_error, updated_at
FROM shopee_fulfillment WHERE sync_status = 'pending' ORDER BY updated_at;
```

### `orders` (Monthly Partitioned)

Order table — `PARTITION BY RANGE (created_at)`, ชื่อ partition `orders_YYYY_MM`

```sql
-- Global sequence (BIGSERIAL per-partition จะมี duplicate id ข้าม partition)
CREATE SEQUENCE IF NOT EXISTS orders_id_seq;

-- Partitioned parent
CREATE TABLE IF NOT EXISTS orders (
    id               BIGINT        NOT NULL DEFAULT nextval('orders_id_seq'),
    shop_id          BIGINT        NOT NULL,
    marketplace_id   TEXT          NOT NULL,   -- order_sn
    marketplace_type TEXT          NOT NULL,   -- "shopee"
    status           TEXT          NOT NULL,
    total_amount     NUMERIC(18,2) NOT NULL DEFAULT 0,
    created_at       TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id, created_at),              -- partition key ต้องอยู่ใน PK
    UNIQUE      (marketplace_id, marketplace_type, created_at)
) PARTITION BY RANGE (created_at);

-- Cross-partition lookup index
CREATE INDEX IF NOT EXISTS orders_marketplace_idx
    ON orders (marketplace_id, marketplace_type);

-- Default catch-all partition
CREATE TABLE IF NOT EXISTS orders_default PARTITION OF orders DEFAULT;

-- Monthly child partitions (EnsureSchema สร้าง current + 2 months ahead)
CREATE TABLE IF NOT EXISTS orders_2026_07
    PARTITION OF orders FOR VALUES FROM ('2026-07-01') TO ('2026-08-01');
```

**Upsert strategy:** PostgreSQL ไม่รองรับ `ON CONFLICT` ข้าม partition → ใช้ find-or-create:
1. `SELECT id, created_at WHERE marketplace_id=$1 AND marketplace_type=$2` (hits index)
2. `ErrNoRows` → INSERT (row ลง partition เดือน `created_at` อัตโนมัติ)
3. Found → `UPDATE ... WHERE id=$3 AND created_at=$4` (partition pruning)

**Migration:** ถ้า `orders` เป็น heap table อยู่แล้ว `EnsureSchema()` rename เป็น `orders_legacy` ก่อน

**Partition creation:** `EnsureSchema()` (startup) สร้าง current month + 2 months ahead. `EnsureMonthPartition()` สามารถเรียกจาก monthly cron เพื่อสร้าง partition ล่วงหน้า.

```sql
-- ดู partition structure
\d+ orders

-- ดู rows per partition
SELECT tableoid::regclass AS partition, count(*)
FROM orders GROUP BY 1 ORDER BY 1;
```

### `shopee_poll_cursor`

Persistent cursor ของ poll job — ป้องกัน full-window scan ทุก restart

```sql
CREATE TABLE shopee_poll_cursor (
    shop_id    BIGINT      PRIMARY KEY,
    cursor_at  TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

### `safety_net_scan_results`

ประวัติ safety-net scan ต่อ run

```sql
CREATE TABLE safety_net_scan_results (
    id               BIGSERIAL   PRIMARY KEY,
    channel          TEXT        NOT NULL,
    shop_id          BIGINT      NOT NULL,
    window_from      TIMESTAMPTZ NOT NULL,
    window_to        TIMESTAMPTZ NOT NULL,
    total_count      INT         NOT NULL DEFAULT 0,
    matched_count    INT         NOT NULL DEFAULT 0,
    missed_count     INT         NOT NULL DEFAULT 0,
    reinjected_count INT         NOT NULL DEFAULT 0,
    duration_ms      BIGINT      NOT NULL DEFAULT 0,
    scanned_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

**Schema สร้างอัตโนมัติตอน startup** ผ่าน `EnsureSchema()` — ไม่ต้องรัน migration มือ

---

## Kafka Topics

| Topic | Key | Producer | Consumer | หมายเหตุ |
|---|---|---|---|---|
| `order.raw.accepted.v1` | `order_sn` | intake.Accept (webhook + poll + scan re-inject) | Classifier | lightweight RawEvent — shop_id + order_sn + status + push_code |
| `order.ingest.shopee.v1` | `order_sn` | Classifier | OrderConsumer (enricher) | canonical IngestEvent with EventType |
| `order.enriched.v1` | `order_sn` | OrderConsumer | IngestionConsumer | full order detail from GetOrderDetail |
| `order.enriched.shopee.dlq.v1` | `order_sn` | OrderConsumer (on failure) | — | DLQ: permanent enrichment failures |
| `order.lifecycle.v1` | `order_sn` | IngestionConsumer | downstream | DB upsert result event |
| `safety-net.scan.complete.v1` | `shop_id` | SafetyNetScanJob | — | governance: scan run metrics per shop |

**ดู messages ใน Redpanda Console:** `http://localhost:8081`

**order.raw.accepted.v1 payload:**

```json
{
  "shop_id": 225997847,
  "order_sn": "26070985HRXASU",
  "status": "READY_TO_SHIP",
  "code": 3,
  "timestamp": 1783421265
}
```

**order.ingest.shopee.v1 payload:**

```json
{
  "channel": "shopee",
  "event_type": "ORDER_READY_TO_SHIP",
  "shop_id": 225997847,
  "order_sn": "26070985HRXASU",
  "status": "READY_TO_SHIP",
  "code": 3,
  "timestamp": 1783421265
}
```

---

## Idempotency Layers

### Layer 1: Webhook dedup (via intake)

```
Redis SetNX  shopee:webhook:dedup:{order_sn}:{timestamp}  EX 24h
```

Shopee retry webhook ด้วย `timestamp` เดิมถ้าไม่ได้รับ 2xx → key เดิม → skip

### Layer 2: Poll dedup (via intake)

```
Redis SetNX  shopee:poll:seen:{order_sn}  EX 24h
```

ป้องกัน order จาก webhook path ถูก publish ซ้ำโดย poll fallback

### Layer 3: Safety-net cross-reference

```
Redis SADD  safety-net:shopee:processed  {order_sn}  (Set TTL 30m)
```

- `intake.Accept()` ทำ SADD หลังทุก successful publish (ทั้ง webhook และ poll path)
- `OrderPollJob` ทำ SMISMEMBER ก่อน intake เพื่อ skip webhook-handled orders เร็ว
- `SafetyNetScanJob` ทำ SMISMEMBER เพื่อหา missed orders แล้ว re-inject

### Layer 4: Fulfillment dedup

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
make run-consumer                   # Kafka consumers (classifier + enricher + ingestion)
make run-scheduler                  # Order poll + safety-net scan scheduler
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
| `UNPAID` | รอชำระ | รับ webhook → intake → classify → enrich (ไม่ enqueue fulfillment) |
| `READY_TO_SHIP` | ชำระแล้ว รอ ship | OrderConsumer enqueue fulfillment |
| `PROCESSED` | ship_order สำเร็จ | fulfillment MarkShipped |
| `SHIPPED` | courier รับพัสดุแล้ว | — |
| `COMPLETED` | ผู้ซื้อได้รับสินค้า | — |
| `CANCELLED` | ยกเลิก | — |
