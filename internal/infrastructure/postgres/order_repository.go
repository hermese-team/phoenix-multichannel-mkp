package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/okdev/marketplace-sync/internal"
	"github.com/okdev/marketplace-sync/internal/domain/order"
	applog "github.com/okdev/marketplace-sync/pkg/logger"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// EnsureSchema creates the orders table as a monthly-partitioned table.
//
// Partition strategy: RANGE on created_at, one child table per calendar month
// named orders_YYYY_MM.  A DEFAULT catch-all partition (orders_default) is
// always present so rows never get rejected before the next month's partition
// is created.
//
// If a non-partitioned orders table already exists (legacy) it is renamed to
// orders_legacy so historical rows are preserved.  Safe to call on every
// startup — all statements are idempotent.
func (r *OrderRepository) EnsureSchema(ctx context.Context) error {
	// 1. Global sequence — survives across partitions (BIGSERIAL would be
	//    per-partition, producing duplicate ids across child tables).
	if _, err := r.db.ExecContext(ctx,
		`CREATE SEQUENCE IF NOT EXISTS orders_id_seq`); err != nil {
		return fmt.Errorf("create orders_id_seq: %w", err)
	}

	// 2. Rename non-partitioned table to orders_legacy if it exists.
	//    pg_class.relkind = 'r' = ordinary (heap) table;
	//                       'p' = partitioned table.
	if _, err := r.db.ExecContext(ctx, `
		DO $$
		BEGIN
			IF EXISTS (
				SELECT 1 FROM pg_class c
				JOIN pg_namespace n ON n.oid = c.relnamespace
				WHERE c.relname = 'orders'
				  AND n.nspname = 'public'
				  AND c.relkind = 'r'
			) THEN
				ALTER TABLE orders RENAME TO orders_legacy;
			END IF;
		END $$`); err != nil {
		return fmt.Errorf("migrate orders_legacy: %w", err)
	}

	// 3. Partitioned parent table.
	//    PRIMARY KEY must include the partition key (PostgreSQL requirement).
	//    UNIQUE likewise includes created_at; since created_at is immutable
	//    after INSERT this is equivalent to the original (marketplace_id,
	//    marketplace_type) uniqueness guarantee.
	if _, err := r.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS orders (
			id               BIGINT        NOT NULL DEFAULT nextval('orders_id_seq'),
			shop_id          BIGINT        NOT NULL,
			marketplace_id   TEXT          NOT NULL,
			marketplace_type TEXT          NOT NULL,
			status           TEXT          NOT NULL,
			total_amount     NUMERIC(18,2) NOT NULL DEFAULT 0,
			created_at       TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
			updated_at       TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
			PRIMARY KEY (id, created_at),
			UNIQUE      (marketplace_id, marketplace_type, created_at)
		) PARTITION BY RANGE (created_at)`); err != nil {
		return fmt.Errorf("create orders partitioned table: %w", err)
	}

	// 4. Non-unique index for cross-partition lookups by business key.
	//    Used by the find-or-create Upsert path to locate existing rows
	//    without knowing which month partition they live in.
	if _, err := r.db.ExecContext(ctx, `
		CREATE INDEX IF NOT EXISTS orders_marketplace_idx
			ON orders (marketplace_id, marketplace_type)`); err != nil {
		return fmt.Errorf("create orders_marketplace_idx: %w", err)
	}

	// 5. DEFAULT catch-all partition — must exist before monthly partitions
	//    so INSERT never fails while a future month's partition is pending.
	if _, err := r.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS orders_default
			PARTITION OF orders DEFAULT`); err != nil {
		return fmt.Errorf("create orders_default partition: %w", err)
	}

	// 6. Create current month + next two months' partitions so the scheduler
	//    doesn't need to run before the first INSERT of a new month.
	now := time.Now()
	for i := 0; i < 3; i++ {
		t := now.AddDate(0, i, 0)
		if err := r.ensureMonthPartition(ctx, t.Year(), int(t.Month())); err != nil {
			return err
		}
	}

	return nil
}

// EnsureMonthPartition creates the child partition for the given year/month if
// it does not already exist.  Call this from a monthly cron job so the next
// month's partition is ready before the first INSERT of that month arrives.
func (r *OrderRepository) EnsureMonthPartition(ctx context.Context, year, month int) error {
	return r.ensureMonthPartition(ctx, year, month)
}

func (r *OrderRepository) ensureMonthPartition(ctx context.Context, year, month int) error {
	name := fmt.Sprintf("orders_%04d_%02d", year, month)
	from := fmt.Sprintf("%04d-%02d-01", year, month)
	nextYear, nextMonth := year, month+1
	if nextMonth > 12 {
		nextYear++
		nextMonth = 1
	}
	to := fmt.Sprintf("%04d-%02d-01", nextYear, nextMonth)

	q := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s
			PARTITION OF orders
			FOR VALUES FROM ('%s') TO ('%s')`,
		name, from, to)

	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("create partition %s: %w", name, err)
	}
	return nil
}

// Upsert inserts a new order or updates an existing one idempotently.
//
// PostgreSQL partitioned tables do not support ON CONFLICT across partitions
// (the partition key must be included in the conflict target, but we don't
// know which month an existing row lives in without querying first).
//
// Strategy: SELECT the existing row to get its (id, created_at), then
// UPDATE; if the row does not exist, INSERT.  The SELECT hits the
// orders_marketplace_idx covering index and touches exactly one partition.
func (r *OrderRepository) Upsert(ctx context.Context, o *order.Order) (isNew bool, err error) {
	var existingID int64
	var existingCreatedAt time.Time

	err = r.db.QueryRowContext(ctx,
		`SELECT id, created_at FROM orders
		  WHERE marketplace_id = $1 AND marketplace_type = $2
		  LIMIT 1`,
		o.SellChannelID, o.SellChannelType,
	).Scan(&existingID, &existingCreatedAt)

	if err != nil && err != sql.ErrNoRows {
		return false, fmt.Errorf("upsert select: %w", err)
	}

	if err == sql.ErrNoRows {
		// INSERT — row goes into the correct monthly partition automatically.
		err = r.db.QueryRowContext(ctx, `
			INSERT INTO orders
				(shop_id, marketplace_id, marketplace_type, status, total_amount, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
			RETURNING id, created_at, updated_at`,
			o.ShopID, o.SellChannelID, o.SellChannelType, string(o.Status), o.TotalAmount,
		).Scan(&o.ID, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return false, fmt.Errorf("upsert insert: %w", err)
		}
		return true, nil
	}

	// UPDATE — supply created_at so Postgres routes to the correct partition.
	o.ID = existingID
	o.CreatedAt = existingCreatedAt
	_, err = r.db.ExecContext(ctx, `
		UPDATE orders
		   SET status       = $1,
		       total_amount = $2,
		       updated_at   = NOW()
		 WHERE id = $3 AND created_at = $4`,
		string(o.Status), o.TotalAmount, existingID, existingCreatedAt,
	)
	if err != nil {
		return false, fmt.Errorf("upsert update: %w", err)
	}
	o.UpdatedAt = time.Now()
	return false, nil
}

func (r *OrderRepository) FindBySellChannelID(ctx context.Context, marketplaceID, marketplaceType string) (*order.Order, error) {
	const q = `
		SELECT id, shop_id, marketplace_id, marketplace_type, status, total_amount, created_at, updated_at
		  FROM orders
		 WHERE marketplace_id = $1 AND marketplace_type = $2
		 LIMIT 1`
	o := &order.Order{}
	err := r.db.QueryRowContext(ctx, q, marketplaceID, marketplaceType).Scan(
		&o.ID, &o.ShopID, &o.SellChannelID, &o.SellChannelType,
		&o.Status, &o.TotalAmount, &o.CreatedAt, &o.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, internal.ErrNotFound
	}
	if err != nil {
		applog.Error("query order", "error", err)
		return nil, err
	}
	return o, nil
}

func (r *OrderRepository) Save(ctx context.Context, o *order.Order) error {
	const q = `
		INSERT INTO orders (shop_id, marketplace_id, marketplace_type, status, total_amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id`
	if err := r.db.QueryRowContext(ctx, q,
		o.ShopID, o.SellChannelID, o.SellChannelType, o.Status, o.TotalAmount,
	).Scan(&o.ID); err != nil {
		applog.Error("insert order", "error", err)
		return err
	}
	return nil
}

func (r *OrderRepository) Update(ctx context.Context, o *order.Order) error {
	const q = `UPDATE orders SET status=$1, updated_at=NOW() WHERE id=$2 AND created_at=$3`
	if _, err := r.db.ExecContext(ctx, q, o.Status, o.ID, o.CreatedAt); err != nil {
		applog.Error("update order", "error", err)
		return err
	}
	return nil
}
