package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter    "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/intake"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
)

const (
	scanWindow   = 30 * time.Minute
	scanLeaseTTL = 14 * time.Minute
	scanLeaseKey = "shopee:scan:lease"
	scanTopic    = "safety-net.scan.complete.v1"
)

// scanCompleteEvent is the governance event published after each scan run.
type scanCompleteEvent struct {
	Channel         string    `json:"channel"`
	ShopID          int64     `json:"shop_id"`
	WindowFrom      time.Time `json:"window_from"`
	WindowTo        time.Time `json:"window_to"`
	TotalCount      int       `json:"total_count"`
	MatchedCount    int       `json:"matched_count"`
	MissedCount     int       `json:"missed_count"`
	ReinjectedCount int       `json:"reinjected_count"`
	DurationMs      int64     `json:"duration_ms"`
	ScannedAt       time.Time `json:"scanned_at"`
}

// SafetyNetScanJob audits the webhook pipeline by comparing GetOrderList
// results against the safety-net Redis set and re-injecting any missed orders.
//
// It is intentionally separate from OrderPollJob so each concern has its own
// schedule, lease, and metrics.
type SafetyNetScanJob struct {
	shopeeClient *client.Client
	tokens       *tokenstore.Store
	tokenRepo    *pgAdapter.ShopeeTokenRepository
	scanRepo     *pgAdapter.SafetyNetScanRepository
	producer     *kafkaAdapter.Producer
	rdb          *redisAdapter.Client
}

func NewSafetyNetScanJob(
	shopeeClient *client.Client,
	tokens *tokenstore.Store,
	tokenRepo *pgAdapter.ShopeeTokenRepository,
	scanRepo *pgAdapter.SafetyNetScanRepository,
	producer *kafkaAdapter.Producer,
	rdb *redisAdapter.Client,
) *SafetyNetScanJob {
	return &SafetyNetScanJob{
		shopeeClient: shopeeClient,
		tokens:       tokens,
		tokenRepo:    tokenRepo,
		scanRepo:     scanRepo,
		producer:     producer,
		rdb:          rdb,
	}
}

// Run is called by the cron scheduler on each tick.
func (j *SafetyNetScanJob) Run() {
	ctx := context.Background()

	// Distributed lease — only one pod scans at a time.
	if j.rdb != nil {
		acquired, err := j.rdb.SetNX(ctx, scanLeaseKey, "1", scanLeaseTTL)
		if err != nil {
			logger.WarnContext(ctx, "scan lease check failed, proceeding anyway",
				"event", "safety_net_scan.lease.error",
				"error", err,
			)
		} else if !acquired {
			logger.DebugContext(ctx, "scan skipped: lease held by another instance",
				"event", "safety_net_scan.lease.skipped",
			)
			return
		}
		defer j.rdb.Delete(ctx, scanLeaseKey)
	}

	shopIDs, err := j.tokenRepo.FindAllShopIDs(ctx)
	if err != nil {
		logger.ErrorContext(ctx, "find shop ids failed",
			"event", "safety_net_scan.shop_ids.error",
			"error", err,
		)
		return
	}

	now := time.Now()
	windowFrom := now.Add(-scanWindow)

	logger.InfoContext(ctx, "safety-net scan started",
		"event", "safety_net_scan.run.started",
		"shops", len(shopIDs),
		"window_from", windowFrom,
		"window_to", now,
	)

	for _, shopID := range shopIDs {
		j.scanShop(ctx, shopID, windowFrom, now)
	}
}

func (j *SafetyNetScanJob) scanShop(ctx context.Context, shopID int64, windowFrom, windowTo time.Time) {
	start := time.Now()

	accessToken, _, err := j.tokens.Get(ctx, shopID)
	if err != nil || accessToken == "" {
		logger.WarnContext(ctx, "no token for shop",
			"event", "safety_net_scan.token.missing",
			"shop_id", shopID,
		)
		return
	}

	orders, err := j.shopeeClient.GetOrderList(ctx, shopID, accessToken, windowFrom.Unix(), windowTo.Unix())
	if err != nil {
		logger.ErrorContext(ctx, "get order list failed",
			"event", "safety_net_scan.get_order_list.error",
			"shop_id", shopID,
			"error", err,
		)
		return
	}

	totalCount := len(orders)
	matchedCount := 0
	missedCount := 0
	reinjectedCount := 0

	if totalCount > 0 && j.rdb != nil {
		// SMISMEMBER: check which orders the webhook pipeline already handled.
		sns := make([]string, len(orders))
		for i, o := range orders {
			sns[i] = o.OrderSN
		}
		inSafetyNet, err := j.rdb.SMIsMember(ctx, intake.SafetyNetKey, sns...)
		if err != nil {
			logger.WarnContext(ctx, "scan SMISMEMBER failed",
				"event", "safety_net_scan.smismember.error",
				"shop_id", shopID,
				"error", err,
			)
		} else {
			for i, o := range orders {
				if inSafetyNet[i] {
					matchedCount++
					continue
				}
				// Missed by webhook — re-inject into the pipeline.
				missedCount++
				raw := intake.RawEvent{
					ShopID:    shopID,
					OrderSN:   o.OrderSN,
					Status:    o.OrderStatus,
					PushCode:  0,
					Timestamp: time.Now().Unix(),
				}
				msgBytes, _ := json.Marshal(raw)
				if err := j.producer.Publish(ctx, intake.OrderRawTopic, []byte(o.OrderSN), msgBytes); err != nil {
					logger.ErrorContext(ctx, "scan re-inject failed",
						"event", "safety_net_scan.reinject.error",
						"order_sn", o.OrderSN,
						"shop_id", shopID,
						"error", err,
					)
				} else {
					reinjectedCount++
				}
			}
		}
	}

	durationMs := time.Since(start).Milliseconds()
	scannedAt := time.Now()

	logger.InfoContext(ctx, "safety-net scan complete",
		"event", "safety_net_scan.shop.complete",
		"shop_id", shopID,
		"total", totalCount,
		"matched", matchedCount,
		"missed", missedCount,
		"reinjected", reinjectedCount,
		"duration_ms", durationMs,
	)

	// Publish governance event.
	evt := scanCompleteEvent{
		Channel:         "shopee",
		ShopID:          shopID,
		WindowFrom:      windowFrom,
		WindowTo:        windowTo,
		TotalCount:      totalCount,
		MatchedCount:    matchedCount,
		MissedCount:     missedCount,
		ReinjectedCount: reinjectedCount,
		DurationMs:      durationMs,
		ScannedAt:       scannedAt,
	}
	if evtBytes, err := json.Marshal(evt); err == nil {
		_ = j.producer.Publish(ctx, scanTopic, []byte(fmt.Sprintf("%d", shopID)), evtBytes)
	}

	// Persist scan result.
	if j.scanRepo != nil {
		res := pgAdapter.ScanResult{
			Channel:         "shopee",
			ShopID:          shopID,
			WindowFrom:      windowFrom,
			WindowTo:        windowTo,
			TotalCount:      totalCount,
			MatchedCount:    matchedCount,
			MissedCount:     missedCount,
			ReinjectedCount: reinjectedCount,
			DurationMs:      durationMs,
			ScannedAt:       scannedAt,
		}
		if err := j.scanRepo.Save(ctx, res); err != nil {
			logger.WarnContext(ctx, "save scan result failed",
				"event", "safety_net_scan.save.error",
				"shop_id", shopID,
				"error", err,
			)
		}
	}
}
