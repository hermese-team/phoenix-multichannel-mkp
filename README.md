# channel-adapter-shopee

Shopee channel adapter สำหรับ Phoenix OMS รับ order events จาก Shopee แล้ว publish ลง Kafka เพื่อให้ downstream consumers นำไปประมวลผลต่อ

## Overview

Service นี้เป็น **anti-corruption layer** ระหว่าง Shopee Open Platform กับ internal event bus ของ Phoenix ทำหน้าที่สองอย่าง:

**Server** (`cmd/server`) — รับ webhook ที่ Shopee ยิงเข้ามา verify signature แล้ว publish raw payload ลง Kafka ทันที ไม่รอ downstream

**Worker** (`cmd/worker`) — poll Shopee API ทุก interval เพื่อดึง order ที่ update ใหม่ เป็น backup สำหรับกรณี webhook miss หรือ service restart

```
┌──────────────────────────────────────────────────────────────┐
│  Shopee Open Platform                                         │
│                                                               │
│  POST /webhook/shopee ──► Server ──► Kafka                    │
│                                       raw.order.accepted.v1  │
│  GET  /v2/order/* ◄────── Worker ──►                         │
└──────────────────────────────────────────────────────────────┘
         │ (both paths)
         ▼
   Redis (token store + rate limit + poller cursor)
```

### Data flow

1. Shopee ยิง push notification มาที่ `POST /webhook/shopee`
2. Server verify HMAC signature แล้ว produce raw JSON ลง `raw.order.accepted.v1`
3. Worker poll `get_order_list` → `get_order_detail` ทุก 30 วินาที produce ลง topic เดียวกัน
4. Downstream (phoenix-oms-mkp-service) consume จาก Kafka แปลงเป็น internal order model

## Architecture

```
cmd/
  server/   HTTP server (webhook + OAuth callback)
  worker/   Polling loop
  authbot/  One-shot tool สำหรับ manual OAuth flow

internal/
  shopee/
    auth.go       OAuth token management (Redis-backed, auto-refresh)
    client.go     Base HTTP client (signed + OTel-traced)
    order.go      Order API (get_order_list, get_order_detail)
    signer.go     HMAC-SHA256 request signing
    model.go      API request/response types

  infrastructure/
    kafka.go      KafkaProducer (franz-go, acks=all, OTel spans)
    redis.go      Redis client (standalone / cluster)
    tracer.go     OTel trace provider (OTLP gRPC)
    logger.go     Zap logger + OTel log bridge

  worker/
    poller.go     Cursor-based polling loop with Redis rate limiter

deployments/
  docker-compose.yml          Redis, Kafka, Redpanda Console, RedisInsight
  pours/deployment/           SigNoz (Foundry-generated)
```

### Key design decisions

- **Kafka เป็น boundary** — server และ worker ไม่รอ OMS ทำงาน publish แล้วตอบ 200 ทันที
- **Token ใน Redis** — ทุก pod ใช้ token เดียวกัน refresh ครั้งเดียวไม่ race กัน
- **Rate limit แบบ distributed** — GCRA limiter ผ่าน Redis ใช้ 80/100 req/min ของ Shopee quota
- **Cursor persisted in Redis** — worker restart ไม่ fetch ซ้ำ overlap-safe ด้วย lookback 5 นาที

## Getting Started

### Prerequisites

- Go 1.22+
- Docker + Docker Compose
- `golangci-lint` (optional, สำหรับ lint)

### 1. Copy config

```bash
cp config.example.json config.json
```

แก้ค่าใน `config.json` หรือ export env var ตาม section [Configuration](#configuration)

### 2. Start local dependencies

```bash
make up
# หรือ
docker compose -f deployments/docker-compose.yml up -d
```

| Service    | URL                   |
|------------|-----------------------|
| Kafka      | `localhost:9092`      |
| Redis      | `localhost:6379`      |
| Kafka UI   | http://localhost:8081 |
| Redis UI   | http://localhost:5540 |

### 3. Run

```bash
# Server (webhook + auth)
make run-server

# Worker (order poller)
make run-worker
```

### OAuth setup (ครั้งแรก)

1. เปิด http://localhost:8080/auth/shopee — browser จะ redirect ไป Shopee consent page
2. Authorize แล้ว Shopee จะ redirect กลับมาที่ `/auth/shopee/callback`
3. Tokens จะถูกเก็บใน Redis อัตโนมัติ worker พร้อมทำงานทันที

สำหรับ sandbox ที่ต้องการ bypass bot detection ให้ใช้ `make auth-bot`

## Configuration

Config โหลดจาก `config.json` แล้ว override ด้วย environment variables

| Field | Env var | Description |
|-------|---------|-------------|
| `shopee.partner_id` | `SHOPEE_PARTNER_ID` | Partner ID จาก Shopee Developer Portal |
| `shopee.partner_key` | `SHOPEE_PARTNER_KEY` | Partner secret key |
| `shopee.shop_id` | `SHOPEE_SHOP_ID` | Shop ID ที่ต้องการ connect |
| `shopee.access_token` | `SHOPEE_ACCESS_TOKEN` | Initial access token (seed ลง Redis ตอน start) |
| `shopee.refresh_token` | `SHOPEE_REFRESH_TOKEN` | Initial refresh token |
| `shopee.base_url` | `SHOPEE_BASE_URL` | `https://openplatform.shopee.sg` (prod) หรือ sandbox URL |
| `kafka.brokers` | — | Kafka broker addresses |
| `kafka.topic_raw_order_accepted` | — | Topic สำหรับ publish orders (default: `raw.order.accepted.v1`) |
| `redis.addrs` | — | Redis addresses |
| `telemetry.endpoint` | — | OTLP gRPC endpoint เช่น `localhost:4317` (เว้นว่างเพื่อ disable) |
| `telemetry.service_name` | — | Service name ที่แสดงใน SigNoz |
| `poller.interval_sec` | — | Polling interval (default: 30) |
| `poller.rate_limit_per_min` | — | Max Shopee API requests/min (default: 80) |
| `poller.lookback_sec` | — | Time window ย้อนหลัง (default: 86400 = 24h) |

## API Reference

### `GET /system/health`

Liveness probe

```json
{ "status": "ok" }
```

### `GET /auth/shopee`

Redirect browser ไปที่ Shopee OAuth consent page สำหรับ initial authorization

### `GET /auth/shopee/callback`

OAuth callback — รับ `code` จาก Shopee แล้ว exchange เป็น access/refresh token เก็บใน Redis

Query params: `code` (string)

```json
{ "status": "ok", "message": "tokens stored — worker is ready" }
```

### `POST /webhook/shopee`

Shopee push notification endpoint ต้อง configure URL นี้ใน Shopee Developer Portal

Request: Shopee-signed JSON payload  
Response: `200 OK` (ตอบทันที ไม่รอ Kafka)

Shopee ส่ง event หลายประเภทมาที่ endpoint นี้ (order status change, logistics update ฯลฯ) service verify HMAC signature แล้ว publish raw payload ลง Kafka ทุก event

## Observability

Service ส่ง traces และ logs ไปที่ SigNoz ผ่าน OpenTelemetry OTLP

### Start SigNoz (local)

```bash
cd deployments/pours/deployment
docker compose up -d
```

SigNoz UI: http://localhost:8090

### Traces

| Span | Description |
|------|-------------|
| `shopee-adapter` (root) | HTTP request เข้า server (otelgin) |
| `shopee GET /api/v2/...` | ทุก Shopee API call |
| `poller.cycle` | รอบ polling แต่ละรอบ — มี attribute `poller.orders_fetched` |
| `kafka.produce` | ทุกครั้งที่ publish ไป Kafka |

### Logs

Logs bridge จาก zap ไป SigNoz ผ่าน otelzap filter ใน Logs Explorer ด้วย:

```
service.name = channel-adapter-shopee
```

### Config

```json
"telemetry": {
  "endpoint": "localhost:4317",
  "service_name": "channel-adapter-shopee"
}
```

เว้น `endpoint` ว่างเพื่อ disable observability (ใช้ no-op provider แทน)

## Development

```bash
make test    # go test ./... -race
make lint    # golangci-lint run ./...
make build   # ./build/server, ./build/worker
```

### Adding a new Shopee API

1. เพิ่ม request/response types ใน `internal/shopee/model.go`
2. เพิ่ม method ใน `internal/shopee/order.go` โดยใช้ `client.get()` / `client.post()` — OTel span inject อัตโนมัติ
3. ถ้าต้อง produce ลง Kafka ให้เรียก `producer.Produce()` — span `kafka.produce` trace อัตโนมัติ

## Dependencies

| Package | Purpose |
|---------|---------|
| `github.com/twmb/franz-go` | Kafka client |
| `github.com/redis/go-redis/v9` | Redis client |
| `github.com/gin-gonic/gin` | HTTP router |
| `github.com/spf13/viper` | Config management |
| `go.uber.org/zap` | Structured logging |
| `go.opentelemetry.io/otel` | OpenTelemetry traces |
| `go.opentelemetry.io/contrib/bridges/otelzap` | Zap → OTLP log bridge |
| `github.com/go-redis/redis_rate/v10` | Distributed rate limiter (GCRA) |
