package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// Seller-voucher create states.
const (
	VoucherPending = "pending"
	VoucherCreated = "created"
	VoucherFailed  = "failed"
)

// LazadaSellerVoucher is a mock row describing a seller voucher we want to create
// on Lazada. Times are epoch milliseconds; discount amounts are stored as text
// to pass through to the API unchanged (Lazada expects decimal strings).
type LazadaSellerVoucher struct {
	VoucherName       string
	VoucherType       string
	Apply             string
	DisplayArea       string
	DiscountType      string
	CriteriaOverMoney string
	PeriodStartTime   int64
	PeriodEndTime     int64
	CollectStartTime  int64
	PerCustomerLimit  int
	Issued            int

	OfferingMoneyValueOff         string
	OfferingPercentageDiscountOff int
	MaxDiscountOfferingMoneyValue string
}

// LazadaSellerVoucherRepository holds vouchers pending creation on Lazada.
type LazadaSellerVoucherRepository struct {
	db *sql.DB
}

func NewLazadaSellerVoucherRepository(db *sql.DB) *LazadaSellerVoucherRepository {
	return &LazadaSellerVoucherRepository{db: db}
}

// EnsureSchema creates the mock table if it does not exist (test convenience;
// a real deployment would use a migration).
func (r *LazadaSellerVoucherRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS lazada_seller_voucher (
		voucher_name          TEXT        PRIMARY KEY,
		voucher_type          TEXT        NOT NULL DEFAULT 'COLLECTIBLE_VOUCHER',
		apply                 TEXT        NOT NULL,
		display_area          TEXT        NOT NULL,
		discount_type         TEXT        NOT NULL,
		criteria_over_money   TEXT        NOT NULL,
		period_start_time     BIGINT      NOT NULL,
		period_end_time       BIGINT      NOT NULL,
		collect_start_time    BIGINT      NOT NULL DEFAULT 0,
		per_customer_limit    INTEGER     NOT NULL,
		issued                INTEGER     NOT NULL,
		offering_money_value_off          TEXT    NOT NULL DEFAULT '',
		offering_percentage_discount_off  INTEGER NOT NULL DEFAULT 0,
		max_discount_offering_money_value TEXT    NOT NULL DEFAULT '',
		sync_status           TEXT        NOT NULL DEFAULT 'pending',
		promotion_id          BIGINT,
		last_error            TEXT        NOT NULL DEFAULT '',
		updated_at            TIMESTAMPTZ NOT NULL DEFAULT now()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure lazada seller-voucher schema: %w", err)
	}
	return nil
}

// ListPending returns up to limit vouchers awaiting creation, oldest first.
func (r *LazadaSellerVoucherRepository) ListPending(ctx context.Context, limit int) ([]LazadaSellerVoucher, error) {
	const q = `SELECT voucher_name, voucher_type, apply, display_area, discount_type,
	                  criteria_over_money, period_start_time, period_end_time, collect_start_time,
	                  per_customer_limit, issued,
	                  offering_money_value_off, offering_percentage_discount_off, max_discount_offering_money_value
	           FROM lazada_seller_voucher
	           WHERE sync_status = $1 ORDER BY updated_at LIMIT $2`
	rows, err := r.db.QueryContext(ctx, q, VoucherPending, limit)
	if err != nil {
		return nil, fmt.Errorf("list pending seller-voucher: %w", err)
	}
	defer rows.Close()

	var out []LazadaSellerVoucher
	for rows.Next() {
		var v LazadaSellerVoucher
		if err := rows.Scan(
			&v.VoucherName, &v.VoucherType, &v.Apply, &v.DisplayArea, &v.DiscountType,
			&v.CriteriaOverMoney, &v.PeriodStartTime, &v.PeriodEndTime, &v.CollectStartTime,
			&v.PerCustomerLimit, &v.Issued,
			&v.OfferingMoneyValueOff, &v.OfferingPercentageDiscountOff, &v.MaxDiscountOfferingMoneyValue,
		); err != nil {
			return nil, fmt.Errorf("scan seller-voucher: %w", err)
		}
		out = append(out, v)
	}
	return out, rows.Err()
}

// MarkCreated records a successful creation and the new Lazada promotion id.
func (r *LazadaSellerVoucherRepository) MarkCreated(ctx context.Context, voucherName string, promotionID int64) error {
	const q = `UPDATE lazada_seller_voucher
	           SET sync_status = $1, promotion_id = $2, last_error = '', updated_at = now()
	           WHERE voucher_name = $3`
	if _, err := r.db.ExecContext(ctx, q, VoucherCreated, promotionID, voucherName); err != nil {
		return fmt.Errorf("mark voucher created: %w", err)
	}
	return nil
}

// MarkFailed records a failed creation attempt with its reason.
func (r *LazadaSellerVoucherRepository) MarkFailed(ctx context.Context, voucherName, errMsg string) error {
	const q = `UPDATE lazada_seller_voucher
	           SET sync_status = $1, last_error = $2, updated_at = now()
	           WHERE voucher_name = $3`
	if _, err := r.db.ExecContext(ctx, q, VoucherFailed, errMsg, voucherName); err != nil {
		return fmt.Errorf("mark voucher failed: %w", err)
	}
	return nil
}
