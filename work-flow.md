# Lazada Sell-Channel — สถานะ Workflow (รันได้แค่ไหน)

เอกสารนี้สรุปว่า loop ที่ทำไว้ **รัน workflow ได้จริงแค่ไหน** และ **อะไรยัง block** การใช้งานจริง
(อัปเดตล่าสุด: 2026-07-03)

> **สรุปหนึ่งบรรทัด:** โครงครบทุก loop รัน end-to-end ต่อ Lazada API จริงได้ **ถ้าป้อน input เอง** —
> เป็นระดับ *demo / integration test* ยังไม่ใช่ *production workflow* เพราะ (1) ฝั่ง input เป็น mock table
> ที่ต้อง insert เอง และ (2) มี payload ที่ reconstruct ไว้ ต้อง verify กับ live API ก่อน

---

## 1. Workflow ที่รันได้เลย (ต่อ Lazada API จริง)

| workflow | เส้นทาง | trigger | สถานะ |
|---|---|---|:---:|
| **Order เข้า — webhook** | Lazada → `server` (:8080) → Kafka → `consumer` → GetOrderDetail/Items → save DB | Lazada ยิง webhook | ✅ |
| **Order เข้า — poll** | `order-scheduler` → GetOrderList → save DB | cron `@every 5m` | ✅ |
| **Product สร้างใหม่** | insert pending → `product-create-scheduler` → `/product/create` | cron | ✅ |
| **Product แก้ราคา/สต็อก** | insert pending → `product-scheduler` → `/product/price_quantity/update` | cron | ✅ |
| **Product แก้ attribute** | insert pending → `product-update-scheduler` → `/product/update` | cron | ✅ |
| **Product ปิดการขาย** | insert pending → `product-offsale-scheduler` → deactivate/remove | cron | ✅ |
| **Stock ตาม warehouse** | insert pending → `sellable-stock-scheduler` → `/product/stock/sellable/*` | cron | ✅ |
| **รูปสินค้า** | insert pending → `image-scheduler` → upload + set images (resumable) | cron | ✅ |
| **Catalog reconcile** | `catalog-scheduler` → GetProducts (paginate) → upsert snapshot | cron `@every 1h` | ✅ |
| **Fulfillment** | insert pending order → `fulfillment-scheduler` → pack/RTS (dropship) + own-fleet ship | cron | ✅ |
| **Promotion — Seller Voucher** | insert pending → `voucher-scheduler` → `/promotion/voucher/create` | cron | ✅ |

> **หัวใจ:** outbound loop ทุกตัวเป็น **pending-queue** — อ่าน row สถานะ `pending` จาก table ของตัวเอง →
> ยิงขึ้น Lazada → mark `created`/`failed` (บางตัว resumable เก็บ intermediate state ก่อนก้าวถัดไป)
> ส่วน order/catalog เป็น **inbound** — ดึงเข้ามาเก็บ

---

## 2. Recipe ขั้นต่ำ — ทำให้เห็นมันวิ่ง

### เตรียม (ครั้งเดียว)
```bash
# infra ขึ้น
pg_isready -h localhost -p 5432 && redis-cli ping && nc -z -w2 localhost 9092 && echo ok

# env + seed token (scheduler/consumer ทุกตัวอ่าน token จาก Redis)
cp .env.example .env          # ใส่ LAZADA_APP_KEY / LAZADA_APP_SECRET จริง
set -a; source .env; set +a
make seed-lazada-token REFRESH=<refresh_token> [ACCESS=<access_token>]
```

### ตัวอย่าง 1 — Order เข้า (ครบ webhook → DB)
```bash
make run-lazada-server        # รับ webhook :8080 → publish Kafka
make run-lazada-consumer      # (อีก terminal) consume → ดึง order detail → save DB
# แล้วให้ Lazada ยิง webhook เข้ามา (หรือ POST /webhook/order เองเพื่อทดสอบ)
```

### ตัวอย่าง 2 — Outbound loop (เช่น voucher)
```bash
make run-lazada-voucher-scheduler   # หยิบ pending row รอบแรกทันที
```
```sql
-- insert งานเข้า table แล้ว scheduler จะยิงขึ้น Lazada รอบถัดไป
INSERT INTO lazada_seller_voucher
  (voucher_name, apply, display_area, discount_type, criteria_over_money,
   period_start_time, period_end_time, per_customer_limit, issued, offering_money_value_off)
VALUES
  ('JULY50', 'ENTIRE_SHOP', 'REGULAR_CHANNEL', 'MONEY_VALUE_OFF', '500',
   1751500000000, 1752000000000, 1, 100, '50');
```
> แต่ละ loop มี table + คอลัมน์ของตัวเอง — ดู `internal/infrastructure/postgres/lazada_*_repository.go` (`EnsureSchema`)

### ตามดู log
```bash
mkdir -p logs
make run-lazada-voucher-scheduler 2>&1 | tee logs/voucher.log
tail -f logs/voucher.log | jq 'select(.level=="error")'
```

---

## 3. จุดที่ยัง block "workflow จริง"

| # | เรื่อง | ระดับ | รายละเอียด |
|:--:|---|:--:|---|
| 1 | **ไม่มี source ป้อน input** | 🔴 | outbound loop ทุกตัวเริ่มจาก **mock table ที่ต้อง insert เอง** — ยังไม่มีตัวเชื่อมจากระบบสินค้า/PIM จริง นี่คือช่องว่างใหญ่สุดถ้าจะใช้งานอัตโนมัติจริง |
| 2 | **payload เป็น mock (9 จุด)** | 🟡 | มี `NOTE(mock)` กำกับในโค้ด ต้อง **verify กับ live API** ก่อนเชื่อว่าถูก — ดูรายการข้อ 4 |
| 3 | **order persist เป็น summary** | 🟡 | เก็บแค่ order number / status / payment / item count ไม่ใช่ line item เต็ม (`sellchannel/lazada/ordersync/syncer.go:47`) |
| 4 | **webhook ไม่ verify signature** | 🟡 | ใครยิง `POST /webhook/order` ปลอมก็เข้าได้ (`sellchannel/lazada/server/webhook.go` NOTE ไว้) |
| 5 | **consumer ไม่มี DLQ/retry** | 🟡 | error แล้วไม่ commit offset → re-read วนซ้ำเรื่อยๆ ไม่มีปลายทางเก็บ message เสีย |
| 6 | **token handling ซ้ำ ~9 ที่** | ⚪ | `clientWithStoredToken`/`refreshTokens` copy อยู่ในทุก syncer — ควร extract เป็น helper ร่วม (tech debt ไม่ block การรัน) |

### รายการ payload ที่ต้อง verify (ข้อ 2)
ไฟล์ที่มี `NOTE(mock)`:
- `sellchannel/lazada/client/fulfillment.go` + `fulfillment_dto.go` — pack/RTS request, own-fleet SOF (.NET ticks)
- `sellchannel/lazada/client/product_offsale.go` — DeactivateProduct request body
- `sellchannel/lazada/client/product_image.go` — upload multipart signature + response shape
- `sellchannel/lazada/client/product_catalog_dto.go` — GetProducts sku field (PascalCase)
- `sellchannel/lazada/client/voucher.go` — promotion envelope (success/error_code แยกจาก gateway code)
- `sellchannel/lazada/productupdatesync/syncer.go` — update payload
- `internal/usecase/order/process_webhook.go` — lookup error ถือเป็น "not found" (mock orders table)
- `sellchannel/lazada/server/webhook.go` — log raw body / signature

---

## 4. เส้นทางสู่ production (ลำดับที่แนะนำ)

1. **อุดข้อ 1** — ออกแบบ source ป้อน input (เชื่อมระบบสินค้าจริง → เขียน pending row) ให้ loop ทำงานอัตโนมัติ
2. **อุดข้อ 2** — verify mock payload ทีละตัวกับ live API (เริ่มจาก loop ที่จะใช้ก่อน)
3. **อุดข้อ 3-5** — order line item เต็ม, webhook signature, consumer DLQ
4. เก็บข้อ 6 (refactor token helper) ตอน refactor รอบถัดไป

---

## 5. Flow ที่ยังไม่ได้ทำ (promotions ที่เหลือ)

| กลุ่ม | สถานะ |
|---|:--:|
| Seller Voucher | 🟡 create loop ✅ — ยังไม่มี activate/deactivate/add-remove-sku (full CRUD) |
| Flexicombo | ⬜ |
| Free Shipping | ⬜ |
| Early Bird Price | ⬜ |

> ดู workflow การรัน + env var ของแต่ละ cmd ที่ [`readme-laza.md`](./readme-laza.md)
