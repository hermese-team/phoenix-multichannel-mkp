-- name: CreateOrder :exec
INSERT INTO orders (id, buyer_id, seller_id, channel, status, total_price, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetOrderByID :one
SELECT id, buyer_id, seller_id, channel, status, total_price, created_at, updated_at
FROM orders
WHERE id = $1;

-- name: UpdateOrder :exec
UPDATE orders
SET status      = $2,
    total_price = $3,
    updated_at  = $4
WHERE id = $1;

-- name: ListOrdersByBuyerID :many
SELECT id, buyer_id, seller_id, channel, status, total_price, created_at, updated_at
FROM orders
WHERE buyer_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountOrdersByBuyerID :one
SELECT count(*) FROM orders WHERE buyer_id = $1;

-- name: CreateOrderItem :exec
INSERT INTO order_items (order_id, product_id, name, price, quantity)
VALUES ($1, $2, $3, $4, $5);

-- name: GetOrderItems :many
SELECT product_id, name, price, quantity
FROM order_items
WHERE order_id = $1;
