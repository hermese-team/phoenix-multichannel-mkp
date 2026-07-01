package product

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	"github.com/ascend/phoenix-multichannel-mkp/internal/repository"
	domainerrors "github.com/ascend/phoenix-multichannel-mkp/pkg/errors"
)

const cacheTTL = 5 * time.Minute

// ProductRepository is a Postgres-backed product store with a Redis read-through cache.
type ProductRepository struct {
	pool  *pgxpool.Pool
	cache *redis.Client
}

func NewProductRepository(pool *pgxpool.Pool, cache *redis.Client) *ProductRepository {
	return &ProductRepository{pool: pool, cache: cache}
}

func cacheKey(id string) string { return "product:" + id }

func (r *ProductRepository) Create(ctx context.Context, p *entity.Product) error {
	const q = `
		INSERT INTO products (id, seller_id, name, description, price, stock, status, channels, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.pool.Exec(ctx, q,
		p.ID, p.SellerID, p.Name, p.Description, p.Price, p.Stock, p.Status, channelsToText(p.Channels), p.CreatedAt, p.UpdatedAt,
	)
	if err != nil {
		return err
	}
	r.cacheSet(ctx, p)
	return nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	if p := r.cacheGet(ctx, id); p != nil {
		return p, nil
	}

	const q = `
		SELECT id, seller_id, name, description, price, stock, status, channels, created_at, updated_at
		FROM products WHERE id = $1`
	p, err := scanProduct(r.pool.QueryRow(ctx, q, id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domainerrors.New("NOT_FOUND", "product not found", domainerrors.ErrNotFound)
		}
		return nil, err
	}
	r.cacheSet(ctx, p)
	return p, nil
}

func (r *ProductRepository) Update(ctx context.Context, p *entity.Product) error {
	const q = `
		UPDATE products
		SET seller_id=$2, name=$3, description=$4, price=$5, stock=$6, status=$7, channels=$8, updated_at=$9
		WHERE id=$1`
	tag, err := r.pool.Exec(ctx, q,
		p.ID, p.SellerID, p.Name, p.Description, p.Price, p.Stock, p.Status, channelsToText(p.Channels), p.UpdatedAt,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domainerrors.New("NOT_FOUND", "product not found", domainerrors.ErrNotFound)
	}
	r.cacheSet(ctx, p)
	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	tag, err := r.pool.Exec(ctx, `DELETE FROM products WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domainerrors.New("NOT_FOUND", "product not found", domainerrors.ErrNotFound)
	}
	r.cacheDel(ctx, id)
	return nil
}

func (r *ProductRepository) List(ctx context.Context, filter repository.ProductFilter) ([]*entity.Product, int, error) {
	where := "WHERE 1=1"
	args := []any{}
	idx := 1
	if filter.SellerID != "" {
		where += fmt.Sprintf(" AND seller_id=$%d", idx)
		args = append(args, filter.SellerID)
		idx++
	}
	if filter.Status != "" {
		where += fmt.Sprintf(" AND status=$%d", idx)
		args = append(args, filter.Status)
		idx++
	}
	if filter.Channel != "" {
		where += fmt.Sprintf(" AND $%d = ANY(channels)", idx)
		args = append(args, string(filter.Channel))
		idx++
	}

	var total int
	if err := r.pool.QueryRow(ctx, "SELECT count(*) FROM products "+where, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	page, limit := filter.Page, filter.Limit
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}
	offset := (page - 1) * limit

	q := fmt.Sprintf(`
		SELECT id, seller_id, name, description, price, stock, status, channels, created_at, updated_at
		FROM products %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, where, idx, idx+1)
	args = append(args, limit, offset)

	rows, err := r.pool.Query(ctx, q, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		p, err := scanProduct(rows)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

// ── Redis cache helpers ─────────────────────────────────────────────────────

func (r *ProductRepository) cacheGet(ctx context.Context, id string) *entity.Product {
	if r.cache == nil {
		return nil
	}
	data, err := r.cache.Get(ctx, cacheKey(id)).Bytes()
	if err != nil {
		return nil
	}
	var p entity.Product
	if json.Unmarshal(data, &p) != nil {
		return nil
	}
	return &p
}

func (r *ProductRepository) cacheSet(ctx context.Context, p *entity.Product) {
	if r.cache == nil {
		return
	}
	if data, err := json.Marshal(p); err == nil {
		_ = r.cache.Set(ctx, cacheKey(p.ID), data, cacheTTL).Err()
	}
}

func (r *ProductRepository) cacheDel(ctx context.Context, id string) {
	if r.cache == nil {
		return
	}
	_ = r.cache.Del(ctx, cacheKey(id)).Err()
}

// ── scanning helpers ────────────────────────────────────────────────────────

type rowScanner interface {
	Scan(dest ...any) error
}

func scanProduct(row rowScanner) (*entity.Product, error) {
	var (
		p        entity.Product
		channels []string
	)
	if err := row.Scan(
		&p.ID, &p.SellerID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Status, &channels, &p.CreatedAt, &p.UpdatedAt,
	); err != nil {
		return nil, err
	}
	p.Channels = textToChannels(channels)
	return &p, nil
}

func channelsToText(channels []entity.Channel) []string {
	out := make([]string, len(channels))
	for i, c := range channels {
		out[i] = string(c)
	}
	return out
}

func textToChannels(values []string) []entity.Channel {
	out := make([]entity.Channel, len(values))
	for i, v := range values {
		out[i] = entity.Channel(v)
	}
	return out
}
