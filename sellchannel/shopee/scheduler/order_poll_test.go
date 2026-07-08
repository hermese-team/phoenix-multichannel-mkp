package scheduler

import (
	"context"
	"fmt"
	"testing"
	"time"

	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
)

// newTestRedis connects to the local Redis used by docker-compose.
// Tests are skipped automatically if Redis is unavailable.
func newTestRedis(t *testing.T) *redisAdapter.Client {
	t.Helper()
	rdb, err := redisAdapter.New(redisAdapter.Config{Addr: "localhost:6379"})
	if err != nil {
		t.Skipf("Redis unavailable (%v) — skipping integration test", err)
	}
	return rdb
}

// TestSafetyNetSMIsMember verifies that orders already handled by the webhook
// (present in the safety-net set) are correctly identified via SMISMEMBER.
func TestSafetyNetSMIsMember(t *testing.T) {
	rdb := newTestRedis(t)
	ctx := context.Background()

	testKey := "test:safety-net:shopee:processed"
	t.Cleanup(func() { _ = rdb.Delete(ctx, testKey) })

	// Seed: ORDER_A and ORDER_B processed by webhook, ORDER_C was missed.
	if err := rdb.SAddWithTTL(ctx, testKey, time.Minute, "ORDER_A", "ORDER_B"); err != nil {
		t.Fatalf("SAddWithTTL: %v", err)
	}

	got, err := rdb.SMIsMember(ctx, testKey, "ORDER_A", "ORDER_B", "ORDER_C")
	if err != nil {
		t.Fatalf("SMIsMember: %v", err)
	}
	if !got[0] {
		t.Error("ORDER_A: expected in safety-net")
	}
	if !got[1] {
		t.Error("ORDER_B: expected in safety-net")
	}
	if got[2] {
		t.Error("ORDER_C: should NOT be in safety-net")
	}
	t.Logf("ORDER_A=%v  ORDER_B=%v  ORDER_C=%v", got[0], got[1], got[2])
}

// TestPollSetNXDedup verifies the SetNX poll-level idempotency layer:
// the first call must return isNew=true, subsequent calls must return false.
func TestPollSetNXDedup(t *testing.T) {
	rdb := newTestRedis(t)
	ctx := context.Background()

	key := fmt.Sprintf("shopee:poll:seen:TEST_%d", time.Now().UnixNano())
	t.Cleanup(func() { _ = rdb.Delete(ctx, key) })

	isNew, err := rdb.SetNX(ctx, key, "1", pollDedupTTL)
	if err != nil {
		t.Fatalf("SetNX (first): %v", err)
	}
	if !isNew {
		t.Error("first SetNX: expected isNew=true")
	}

	isNew2, err := rdb.SetNX(ctx, key, "1", pollDedupTTL)
	if err != nil {
		t.Fatalf("SetNX (second): %v", err)
	}
	if isNew2 {
		t.Error("second SetNX: expected isNew=false (duplicate)")
	}
	t.Log("SetNX dedup: first=true, second=false ✓")
}

// TestDistributedLease verifies that the poll lease prevents concurrent runs:
// the second acquire must fail while the first lease is still held.
func TestDistributedLease(t *testing.T) {
	rdb := newTestRedis(t)
	ctx := context.Background()

	testLeaseKey := "test:shopee:poll:lease"
	t.Cleanup(func() { _ = rdb.Delete(ctx, testLeaseKey) })

	// First acquire — should succeed.
	acquired, err := rdb.SetNX(ctx, testLeaseKey, "1", time.Minute)
	if err != nil {
		t.Fatalf("SetNX lease (first): %v", err)
	}
	if !acquired {
		t.Error("first acquire: expected acquired=true")
	}

	// Second acquire — should fail (lease already held).
	acquired2, err := rdb.SetNX(ctx, testLeaseKey, "1", time.Minute)
	if err != nil {
		t.Fatalf("SetNX lease (second): %v", err)
	}
	if acquired2 {
		t.Error("second acquire: expected acquired=false (lease held)")
	}
	t.Log("distributed lease: first=acquired, second=blocked ✓")
}
