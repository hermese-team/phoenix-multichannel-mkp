package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// ScanResult holds the outcome of one safety-net scan run for a single shop.
type ScanResult struct {
	Channel         string
	ShopID          int64
	WindowFrom      time.Time
	WindowTo        time.Time
	TotalCount      int
	MatchedCount    int
	MissedCount     int
	ReinjectedCount int
	DurationMs      int64
	ScannedAt       time.Time
}

// SafetyNetScanRepository persists safety-net scan results for audit and
// capacity-planning dashboards.
type SafetyNetScanRepository struct {
	db *sql.DB
}

func NewSafetyNetScanRepository(db *sql.DB) *SafetyNetScanRepository {
	return &SafetyNetScanRepository{db: db}
}

// EnsureSchema creates the safety_net_scan_results table if it does not exist.
func (r *SafetyNetScanRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS safety_net_scan_results (
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
		scanned_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure safety_net_scan_results schema: %w", err)
	}
	return nil
}

const scanResultRetention = 30 // days

// Save inserts a new scan result row and prunes rows older than retention window.
func (r *SafetyNetScanRepository) Save(ctx context.Context, res ScanResult) error {
	const insert = `
	INSERT INTO safety_net_scan_results
		(channel, shop_id, window_from, window_to,
		 total_count, matched_count, missed_count, reinjected_count,
		 duration_ms, scanned_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	if _, err := r.db.ExecContext(ctx, insert,
		res.Channel, res.ShopID, res.WindowFrom, res.WindowTo,
		res.TotalCount, res.MatchedCount, res.MissedCount, res.ReinjectedCount,
		res.DurationMs, res.ScannedAt,
	); err != nil {
		return fmt.Errorf("save scan result: %w", err)
	}

	// Best-effort cleanup — non-fatal if it fails.
	_, _ = r.db.ExecContext(ctx,
		`DELETE FROM safety_net_scan_results WHERE scanned_at < NOW() - ($1 || ' days')::INTERVAL`,
		scanResultRetention,
	)
	return nil
}
