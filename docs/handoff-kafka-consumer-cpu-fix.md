# Handoff: Kafka Consumer CPU Busy-Loop Fix

> เอกสารส่งต่อ context จากอีก session (ที่ debug Kafka CPU สูงในโปรเจกต์ `kafka-docker`)
> วันที่: 2026-07-02 — อ่านไฟล์นี้เพื่อรับ context ต่อได้เลย

## TL;DR
Kafka container กิน CPU ~43% เพราะ **consumer client (kgo) ของ repo นี้ค้างใน busy-loop**
ต้นเหตุ: consumer loop วน `continue` ตอน fetch error โดย **ไม่มี backoff** → พอเจอ error ต่อเนื่อง
(จาก orphan process หลาย instance ใน group เดียว) loop หมุน ~700 ล้านรอบ/วิ กิน CPU เต็ม core
+ ยิง request ถล่ม broker

## สิ่งที่ session นั้นแก้ในโปรเจกต์นี้ (ยังไม่ได้ commit)

### 1) ใส่ backoff กัน busy-loop — ไฟล์:
- `sellchannel/lazada/consumer/consumer.go`
- `sellchannel/shopee/consumer/product.go`
```go
if errs := fetches.Errors(); len(errs) > 0 {
    for _, e := range errs { logger.Error("kafka fetch", "error", e.Err) }
    // Back off so a persistent fetch error can't spin this loop at full CPU.
    select {
    case <-ctx.Done():
        return nil
    case <-time.After(time.Second):
    }
    continue
}
```
(เพิ่ม import `"time"` ทั้งสองไฟล์)

### 2) Refactor: แยก constructor ตาม role — ไฟล์:
- `sellchannel/lazada/lazada.go` — เดิม `New()` fat constructor สร้างทุกอย่าง →
  แยกเป็น `NewServer()`, `NewScheduler()`, `NewConsumer()` + helper `newSyncer()`
- `sellchannel/shopee/shopee.go` — แยกเป็น `NewServer()`, `NewConsumer()`, `NewScheduler()`
  (แถมแก้ smell: shopee เดิมเปิด Redis แล้ว `_` ทิ้ง — ตัดออก)
- อัปเดต main ทั้ง 6: `sellchannel/{lazada,shopee}/cmd/{server,scheduler,consumer}/main.go`
  (เลิกเรียก `X.New()` + accessor `.Consumer()/.Server()/.Scheduler()` → เรียก constructor ตรง)

**เหตุผล:** consumer เดิมสร้าง kafka producer + http server ทิ้งเปล่า ๆ (fat constructor)
→ แยกให้ role รันเฉพาะ dependency ที่ตัวเองใช้

### 3) Makefile — เพิ่ม target `kill-consumers`
```makefile
kill-consumers: ## Kill orphaned consumer processes left behind by `go run`
	@pkill -f 'exe/consumer'        || true
	@pkill -f 'bin/[a-z]*-consumer' || true
```

## เทสแล้ว ✅
- `go build ./...` + `go vet ./...` ผ่าน
- A/B กลไก loop: OLD ~700M รอบ/วิ (spin), NEW ~1 รอบ/วิ (idle)
- รัน binary จริง 3 instance/group เดียว กับ Kafka จริง: client 0%, kafka ~2%, rebalance 0

## Follow-up ที่ยังไม่ทำ (ถ้าจะทำต่อ)
- **แยกชนิด error ให้ละเอียดขึ้น** (ตอนนี้ backoff เหมาเข่งทุก error): เช่น
  `context.Canceled` → ออกเลย, error retriable → backoff, fatal (unknown topic/auth) → พิจารณา fail fast
- **Commit เป็น batch** แทน `CommitRecords` ต่อ record (lazada) — throughput ดีกว่าถ้าปริมาณโต
- backoff กัน CPU พุ่งได้ แต่ **ไม่ลบสาเหตุ error** — ควรดู log `kafka fetch error` ว่า error จริงคืออะไร

## หมายเหตุ operation
- ต้นตอ orphan มาจาก `go run` ทิ้ง child process เวลา rerun → สะสมหลาย instance ใน group เดียว
- แนะนำ dev ด้วย `go build` เป็น binary (signal ถึงตรง) + `make kill-consumers` ก่อนรันใหม่
- runbook เต็มอยู่ที่โปรเจกต์ `kafka-docker`: `root-cause-infinite-loop-kafka.md`
