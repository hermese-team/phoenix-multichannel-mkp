-- name: CreateProduct :exec
INSERT INTO products (id, seller_id, name, description, price, stock, status, channels, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: GetProductByID :one
SELECT id, seller_id, name, description, price, stock, status, channels, created_at, updated_at
FROM products
WHERE id = $1;

-- name: UpdateProduct :exec
UPDATE products
SET seller_id   = $2,
    name        = $3,
    description = $4,
    price       = $5,
    stock       = $6,
    status      = $7,
    channels    = $8,
    updated_at  = $9
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;

-- name: ListProductsBySellerID :many
SELECT id, seller_id, name, description, price, stock, status, channels, created_at, updated_at
FROM products
WHERE seller_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountProductsBySellerID :one
SELECT count(*) FROM products WHERE seller_id = $1;
