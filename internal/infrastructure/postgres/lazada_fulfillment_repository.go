package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// Fulfillment delivery models + sync states.
const (
	DeliveryDropship = "dropship"  // Lazada logistics: Pack → ReadyToShip
	DeliveryOwnFleet = "own_fleet" // seller ships: push shipped status

	FulfillmentPending     = "pending"       // not started
	FulfillmentPacked      = "packed"        // dropship: packed, awaiting RTS (resumable)
	FulfillmentReadyToShip = "ready_to_ship" // dropship: RTS done (terminal)
	FulfillmentShipped     = "shipped"       // own_fleet: marked shipped (terminal)
)

// LazadaFulfillment is a mock row describing an order (items) to fulfill.
// For own_fleet the caller provides package_id + tracking_number (the seller's
// own shipment); for dropship they are filled in from Pack's response.
type LazadaFulfillment struct {
	ID               int64
	OrderID          int64
	OrderItemIDs     string // comma-separated order_item_id
	DeliveryType     string // DeliveryDropship | DeliveryOwnFleet
	ShipmentProvider string
	PackageID        string
	TrackingNumber   string
	CarrierCode      string
	SyncStatus       string // current progress; lets a dropship row resume at RTS
}

// LazadaFulfillmentRepository holds orders pending fulfillment (test schema).
type LazadaFulfillmentRepository struct {
	db *sql.DB
}

func NewLazadaFulfillmentRepository(db *sql.DB) *LazadaFulfillmentRepository {
	return &LazadaFulfillmentRepository{db: db}
}

// EnsureSchema creates the mock table if it does not exist (test convenience;
// a real deployment would use a migration).
func (r *LazadaFulfillmentRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS lazada_fulfillment (
		id                BIGSERIAL   PRIMARY KEY,
		order_id          BIGINT      NOT NULL,
		order_item_ids    TEXT        NOT NULL,
		delivery_type     TEXT        NOT NULL,
		shipment_provider TEXT        NOT NULL DEFAULT '',
		package_id        TEXT        NOT NULL DEFAULT '',
		tracking_number   TEXT        NOT NULL DEFAULT '',
		carrier_code      TEXT        NOT NULL DEFAULT '',
		sync_status       TEXT        NOT NULL DEFAULT 'pending',
		last_error        TEXT        NOT NULL DEFAULT '',
		updated_at        TIMESTAMPTZ NOT NULL DEFAULT now()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure lazada fulfillment schema: %w", err)
	}
	return nil
}

// ListPending returns up to limit fulfillments still in progress (pending or
// packed-awaiting-RTS), oldest first. Including "packed" lets a dropship row
// resume at ReadyToShip without packing again.
func (r *LazadaFulfillmentRepository) ListPending(ctx context.Context, limit int) ([]LazadaFulfillment, error) {
	const q = `SELECT id, order_id, order_item_ids, delivery_type, shipment_provider,
	                  package_id, tracking_number, carrier_code, sync_status
	           FROM lazada_fulfillment
	           WHERE sync_status IN ($1, $2) ORDER BY updated_at LIMIT $3`
	rows, err := r.db.QueryContext(ctx, q, FulfillmentPending, FulfillmentPacked, limit)
	if err != nil {
		return nil, fmt.Errorf("list pending fulfillment: %w", err)
	}
	defer rows.Close()

	var out []LazadaFulfillment
	for rows.Next() {
		var f LazadaFulfillment
		if err := rows.Scan(
			&f.ID, &f.OrderID, &f.OrderItemIDs, &f.DeliveryType, &f.ShipmentProvider,
			&f.PackageID, &f.TrackingNumber, &f.CarrierCode, &f.SyncStatus,
		); err != nil {
			return nil, fmt.Errorf("scan fulfillment: %w", err)
		}
		out = append(out, f)
	}
	return out, rows.Err()
}

// MarkStatus advances a row to a new progress status, storing the package id +
// tracking number (from Pack, or supplied for own-fleet) and clearing last_error.
func (r *LazadaFulfillmentRepository) MarkStatus(ctx context.Context, id int64, status, packageID, trackingNumber string) error {
	const q = `UPDATE lazada_fulfillment
	           SET sync_status = $1, package_id = $2, tracking_number = $3,
	               last_error = '', updated_at = now()
	           WHERE id = $4`
	if _, err := r.db.ExecContext(ctx, q, status, packageID, trackingNumber, id); err != nil {
		return fmt.Errorf("mark fulfillment status: %w", err)
	}
	return nil
}

// MarkError records a failure without changing the progress status, so the row
// is retried (resumed) on the next run rather than stranded.
func (r *LazadaFulfillmentRepository) MarkError(ctx context.Context, id int64, errMsg string) error {
	const q = `UPDATE lazada_fulfillment SET last_error = $1, updated_at = now() WHERE id = $2`
	if _, err := r.db.ExecContext(ctx, q, errMsg, id); err != nil {
		return fmt.Errorf("mark fulfillment error: %w", err)
	}
	return nil
}
