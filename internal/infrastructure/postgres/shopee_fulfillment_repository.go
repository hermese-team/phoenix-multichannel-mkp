package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// Delivery models + sync states for Shopee fulfillment.
const (
	ShopeeDeliveryShopeeLogistics = "shopee_logistics" // Shopee arranges pickup/dropoff
	ShopeeDeliveryOwnFleet        = "own_fleet"        // seller arranges own shipment

	ShopeeFulfillmentPending = "pending" // not yet shipped — will retry on next tick
	ShopeeFulfillmentShipped = "shipped" // terminal: ship_order called successfully
	ShopeeFulfillmentFailed  = "failed"  // terminal: permanent error, ops intervention required
)

// ShopeeFulfillment represents one order pending fulfillment.
// For own_fleet the caller provides tracking_number; for shopee_logistics it is
// returned by the API after ship_order.
type ShopeeFulfillment struct {
	ID             int64
	ShopID         int64
	OrderSN        string
	DeliveryType   string // ShopeeDeliveryShopeeLogistics | ShopeeDeliveryOwnFleet
	TrackingNumber string // provided (own_fleet) or returned by API (shopee_logistics)
	SyncStatus     string
}

// ShopeeFulfillmentRepository persists orders pending fulfillment.
type ShopeeFulfillmentRepository struct {
	db *sql.DB
}

func NewShopeeFulfillmentRepository(db *sql.DB) *ShopeeFulfillmentRepository {
	return &ShopeeFulfillmentRepository{db: db}
}

// EnsureSchema creates the table if it does not exist.
func (r *ShopeeFulfillmentRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS shopee_fulfillment (
		id              BIGSERIAL   PRIMARY KEY,
		shop_id         BIGINT      NOT NULL,
		order_sn        TEXT        NOT NULL,
		delivery_type   TEXT        NOT NULL DEFAULT 'shopee_logistics',
		tracking_number TEXT        NOT NULL DEFAULT '',
		sync_status     TEXT        NOT NULL DEFAULT 'pending',
		last_error      TEXT        NOT NULL DEFAULT '',
		updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
		UNIQUE (shop_id, order_sn)
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure shopee_fulfillment schema: %w", err)
	}
	return nil
}

// Enqueue inserts a fulfillment row, ignoring conflicts (idempotent on order_sn).
func (r *ShopeeFulfillmentRepository) Enqueue(ctx context.Context, shopID int64, orderSN, deliveryType, trackingNumber string) error {
	const q = `
	INSERT INTO shopee_fulfillment (shop_id, order_sn, delivery_type, tracking_number)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (shop_id, order_sn) DO NOTHING`
	if _, err := r.db.ExecContext(ctx, q, shopID, orderSN, deliveryType, trackingNumber); err != nil {
		return fmt.Errorf("enqueue shopee fulfillment: %w", err)
	}
	return nil
}

// ListPending returns up to limit pending fulfillments, oldest first.
func (r *ShopeeFulfillmentRepository) ListPending(ctx context.Context, limit int) ([]ShopeeFulfillment, error) {
	const q = `
	SELECT id, shop_id, order_sn, delivery_type, tracking_number, sync_status
	FROM shopee_fulfillment
	WHERE sync_status = $1
	ORDER BY updated_at
	LIMIT $2`
	rows, err := r.db.QueryContext(ctx, q, ShopeeFulfillmentPending, limit)
	if err != nil {
		return nil, fmt.Errorf("list pending shopee fulfillment: %w", err)
	}
	defer rows.Close()
	var out []ShopeeFulfillment
	for rows.Next() {
		var f ShopeeFulfillment
		if err := rows.Scan(&f.ID, &f.ShopID, &f.OrderSN, &f.DeliveryType, &f.TrackingNumber, &f.SyncStatus); err != nil {
			return nil, fmt.Errorf("scan shopee fulfillment: %w", err)
		}
		out = append(out, f)
	}
	return out, rows.Err()
}

// MarkShipped marks a row as shipped and stores the tracking number.
func (r *ShopeeFulfillmentRepository) MarkShipped(ctx context.Context, id int64, trackingNumber string) error {
	const q = `
	UPDATE shopee_fulfillment
	SET sync_status = $1, tracking_number = $2, last_error = '', updated_at = now()
	WHERE id = $3`
	if _, err := r.db.ExecContext(ctx, q, ShopeeFulfillmentShipped, trackingNumber, id); err != nil {
		return fmt.Errorf("mark shopee fulfillment shipped: %w", err)
	}
	return nil
}

// MarkError records an error without changing the status so the row is retried.
func (r *ShopeeFulfillmentRepository) MarkError(ctx context.Context, id int64, errMsg string) error {
	const q = `UPDATE shopee_fulfillment SET last_error = $1, updated_at = now() WHERE id = $2`
	if _, err := r.db.ExecContext(ctx, q, errMsg, id); err != nil {
		return fmt.Errorf("mark shopee fulfillment error: %w", err)
	}
	return nil
}

// MarkFailed marks a row as permanently failed (ops intervention required).
// Unlike MarkError, this sets sync_status = 'failed' so the scheduler skips it.
func (r *ShopeeFulfillmentRepository) MarkFailed(ctx context.Context, id int64, errMsg string) error {
	const q = `
	UPDATE shopee_fulfillment
	SET sync_status = $1, last_error = $2, updated_at = now()
	WHERE id = $3`
	if _, err := r.db.ExecContext(ctx, q, ShopeeFulfillmentFailed, errMsg, id); err != nil {
		return fmt.Errorf("mark shopee fulfillment failed: %w", err)
	}
	return nil
}
