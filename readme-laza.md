# Lazada Sell-Channel — คู่มือการรัน

ช่องทาง Lazada ประกอบด้วย **5 binary** แต่ละตัวรันแยกกันและใช้ dependency เฉพาะที่ตัวเองต้องการ
(ตาม design ใน `lazada.go`: consumer ไม่เปิด Kafka producer, scheduler ไม่เปิด HTTP server ฯลฯ)

---

## สรุปแต่ละ cmd — port / dependency

| cmd | หน้าที่ | HTTP Port | Postgres | Redis | Kafka | ต้องมี token ใน Redis |
|---|---|:---:|:---:|:---:|:---:|:---:|
| `server` | รับ webhook จาก Lazada → เช็ค idem (Redis) → publish เข้า Kafka | **APP_PORT (8080)** | ✅ | ✅ | ✅ producer | – |
| `order-scheduler` | poll order list ตาม cron → sync ลง DB | – (ไม่ bind port) | ✅ | ✅ | – | ✅ |
| `product-scheduler` | push price/stock ของ product (pending) ขึ้น Lazada ตาม cron | – (ไม่ bind port) | ✅ | ✅ | – | ✅ |
| `product-create-scheduler` | สร้าง product ใหม่ (pending) บน Lazada ตาม cron | – (ไม่ bind port) | ✅ | ✅ | – | ✅ |
| `consumer` | อ่าน webhook จาก Kafka → ดึง order detail/items → ลง DB | – (ไม่ bind port) | ✅ | ✅ | ✅ consumer group | ✅ |
| `lazada-token` | seed access/refresh token ลง Redis (รันครั้งเดียว) | – | – | ✅ | – | – (เป็นตัวเขียน token) |
| `lazada-ping` | smoke test — เช็ค auth/signing/token กับ Lazada API | – | – | ✅ | – | ✅ |
| `lazada-category` | browse category tree / attributes (หา category_id + attribute ก่อน seed create) | – | – | – | – | – (no-auth) |

> **มีตัวเดียวที่ใช้ port** คือ `server` → ฟังที่ `:APP_PORT` (default **8080**) เปิด `GET /health` และ `POST /webhook/order`
> ตัวอื่นเป็น background worker ไม่เปิด port

---

## Prerequisites

### 1. Infra (Docker local)
| บริการ | Address | หมายเหตุ |
|---|---|---|
| Postgres | `localhost:5432` | user/pass `admin/admin`, db `marketplace_sync` (schema `lazada_synced_orders` สร้างอัตโนมัติตอน scheduler/consumer เริ่ม) |
| Redis | `localhost:6379` | เก็บ token + webhook idempotency key |
| Kafka | `localhost:9092` | topic `lazada.order.webhook` |

ตรวจว่า infra พร้อม:
```bash
pg_isready -h localhost -p 5432
redis-cli ping
nc -z -w2 localhost 9092 && echo "kafka ok"
```

### 2. Environment
```bash
cp .env.example .env      # ใส่ค่าจริงของ LAZADA_APP_KEY / LAZADA_APP_SECRET
set -a; source .env; set +a
```

> ⚠️ `.env` ถูก gitignore — มี app secret จริง ห้าม commit
> access/refresh token **ไม่ได้อยู่ใน .env** แต่อยู่ใน Redis (seed ด้วย `lazada-token`)

---

## ลำดับการรัน

### Step 0 — seed token ลง Redis (ครั้งแรกครั้งเดียว)
`scheduler`, `consumer`, `lazada-ping` ทุกตัวอ่าน token จาก Redis — ต้อง seed ก่อน ไม่งั้น start ไม่ผ่าน
```bash
make seed-lazada-token REFRESH=<refresh_token> [ACCESS=<access_token>]
# หรือ
go run ./sellchannel/lazada/cmd/lazada-token -refresh=<refresh_token> [-access=<access_token>]
```
เช็คว่าเข้าไปแล้ว:
```bash
redis-cli GET lazada:access_token
redis-cli GET lazada:refresh_token
```

### Step 1 — smoke test (แนะนำก่อนรันจริง)
```bash
go run ./sellchannel/lazada/cmd/lazada-ping
# ทดสอบดึง order ด้วย:  LAZADA_ORDER_ID=xxxx go run ./sellchannel/lazada/cmd/lazada-ping
```
ถ้าเห็น `OK /seller/get` = auth/signing/token ใช้ได้

### Step 2 — รัน worker ที่ต้องการ

**Server (webhook receiver):**
```bash
make run-lazada-server        # go run
# หรือ binary:
go build -o bin/lazada-server ./sellchannel/lazada/cmd/server
./bin/lazada-server
```
ฟังที่ `:8080` → เช็ค `curl localhost:8080/health`

**Order scheduler (poll order list):**
```bash
make run-lazada-order-scheduler
```

**Product scheduler (push price/stock):**
```bash
make run-lazada-product-scheduler
```

**Consumer (kafka → order sync):**
```bash
make run-lazada-consumer
```

> **server กับ consumer เป็นคนละหน้าที่** — server แค่รับ webhook แล้วโยนเข้า Kafka,
> ต้องมี consumer รันคู่ถึงจะมีคนดึง order detail มาลง DB จริง

---

## Build เป็น binary ทั้งหมด
```bash
make build     # สร้าง bin/lazada-{server,order-scheduler,product-scheduler,consumer,token} (+ shopee)
```
> แนะนำรันด้วย binary มากกว่า `go run` เวลา dev worker — signal (Ctrl+C) ส่งถึง process ตรง
> ไม่เหลือ orphan (`go run` ทิ้ง child process ค้างได้)

เช็ค/เก็บกวาด consumer ค้าง:
```bash
make check-consumers    # list process ที่ยังรันอยู่
make kill-consumers     # kill orphan
```

---

## Environment Variables

### App
| Variable | Default | ใช้โดย |
|---|---|---|
| `APP_ENV` | `development` | ทุกตัว |
| `APP_PORT` | `8080` | **server** (HTTP port) |
| `LOG_LEVEL` | `info` | ทุกตัว (zap) |

### Lazada API
| Variable | Default | หมายเหตุ |
|---|---|---|
| `LAZADA_APP_KEY` | – | จำเป็น |
| `LAZADA_APP_SECRET` | – | จำเป็น (secret จริง — อย่า commit) |
| `LAZADA_URL_API` | `https://api.lazada.co.th/rest` | API gateway |
| `LAZADA_AUTH_URL` | `https://auth.lazada.com/rest` | endpoint refresh token |
| `LAZADA_ORDER_ID` | – | optional — สำหรับ `lazada-ping` |
| `LAZADA_ACCESS_TOKEN` | – | optional override เฉพาะ `lazada-ping` (ปกติอ่านจาก Redis) |

### Infra
| Variable | Default | ใช้โดย |
|---|---|---|
| `POSTGRES_DSN` | `postgres://postgres:postgres@localhost:5432/marketplace_sync?sslmode=disable` | server, scheduler, consumer |
| `POSTGRES_MAX_OPEN_CONNS` | `25` | |
| `POSTGRES_MAX_IDLE_CONNS` | `5` | |
| `POSTGRES_CONN_MAX_LIFETIME` | `5m` | |
| `REDIS_ADDR` | `localhost:6379` | ทุกตัว |
| `REDIS_PASSWORD` | – | |
| `REDIS_DB` | `0` | |
| `KAFKA_BROKER` | `localhost:9092` | server (producer), consumer |
| `KAFKA_GROUP_ID` | `marketplace-sync` | |

### Webhook (server)
| Variable | Default | หมายเหตุ |
|---|---|---|
| `LAZADA_WEBHOOK_IDEM_PREFIX` | `lazada:webhook:idem` | prefix key idempotency ใน Redis |
| `LAZADA_WEBHOOK_TOPIC` | `lazada.order.webhook` | topic ที่ publish (server) / subscribe (consumer) |

### Consumer
| Variable | Default | |
|---|---|---|
| `LAZADA_WEBHOOK_GROUP` | `lazada-webhook-consumer` | consumer group id |
| `LAZADA_WEBHOOK_TOPIC` | `lazada.order.webhook` | topic ที่ subscribe |

### Scheduler
| Variable | Default | |
|---|---|---|
| `LAZADA_SYNC_SPEC` | `@every 5m` | cron spec |
| `LAZADA_SYNC_WINDOW` | `24h` | ช่วงเวลาย้อนหลังที่ poll |
| `LAZADA_SYNC_LIMIT` | `50` | จำนวน order สูงสุดต่อรอบ |

---

## Data Flow (ภาพรวม)

```
                       ┌─────────────┐
  Lazada  ──webhook──▶ │   server    │ ──publish──▶  Kafka topic
                       │   :8080     │              lazada.order.webhook
                       └─────────────┘                     │
                          │ SETNX idem (Redis)             │ consume (group)
                          ▼                                ▼
                       (dedupe)                     ┌─────────────┐
                                                    │  consumer   │
  Lazada order list ──poll──▶ ┌───────────┐        │             │
                              │ scheduler │──┐      └─────────────┘
                              │  @every5m │  │            │ GetOrderDetail
                              └───────────┘  │            │ GetOrderItems
                                             ▼            ▼
                                          ┌──────────────────┐
                                          │    Postgres      │
                                          │ lazada_synced_.. │
                                          └──────────────────┘
```

## การรัน + เก็บ log (วิธีที่ 1 — `tee` ไฟล์)

zap เขียน log เป็น JSON ออก **stdout** — วิธีที่ 1 คือ redirect ผ่าน `tee` เพื่อ
**เห็นบนจอสด + เก็บลงไฟล์** พร้อมกัน (ไม่ต้องแก้ code) แล้วใช้ `tail -f` / `jq` ตามดูย้อนหลัง

### เตรียม
```bash
mkdir -p logs
make build                       # สร้าง binary ไว้ก่อน (แนะนำมากกว่า go run สำหรับ worker)
set -a; source .env; set +a
```

### รันแต่ละ cmd (เก็บ log แยกไฟล์)
```bash
# server (webhook :8080)
./bin/lazada-server    2>&1 | tee logs/lazada-server.log

# order scheduler
./bin/lazada-order-scheduler 2>&1 | tee logs/lazada-order-scheduler.log

# consumer
./bin/lazada-consumer  2>&1 | tee logs/lazada-consumer.log
```
- `2>&1` = รวม stderr เข้า stdout (กันเหนียว เผื่อ log บางบรรทัดออก stderr)
- `tee logs/xxx.log` = พิมพ์บนจอ **และ** เขียนไฟล์ไปพร้อมกัน
- อยากรัน background ก็ต่อ `&` ท้าย แล้วตามดูด้วย `tail -f` ข้างล่าง

> รันด้วย binary (ไม่ใช่ `go run`) เพื่อให้ Ctrl+C ส่ง signal ถึง process ตรง ไม่เหลือ orphan

### ตามดู log ย้อนหลัง / real-time
```bash
tail -f logs/lazada-consumer.log                              # ไล่ดูสด
tail -f logs/lazada-server.log   | jq 'select(.level=="error")'   # เฉพาะ error
tail -f logs/lazada-consumer.log | jq 'select(.msg=="kafka fetch")'  # เฉพาะ event ที่สนใจ
grep -c '"level":"error"' logs/lazada-order-scheduler.log    # นับ error ในไฟล์
```

> ไฟล์ใน `logs/` ควรใส่ `.gitignore` (อย่า commit log)

