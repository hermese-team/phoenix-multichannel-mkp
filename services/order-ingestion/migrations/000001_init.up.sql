CREATE TABLE IF NOT EXISTS products (
    id          UUID PRIMARY KEY,
    seller_id   TEXT          NOT NULL,
    name        TEXT          NOT NULL,
    description TEXT          NOT NULL DEFAULT '',
    price       NUMERIC(12,2) NOT NULL,
    stock       INTEGER       NOT NULL,
    status      TEXT          NOT NULL,
    channels    TEXT[]        NOT NULL DEFAULT '{}',
    created_at  TIMESTAMPTZ   NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ   NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_products_seller_id ON products (seller_id);
CREATE INDEX IF NOT EXISTS idx_products_status    ON products (status);

CREATE TABLE IF NOT EXISTS orders (
    id          UUID PRIMARY KEY,
    buyer_id    TEXT          NOT NULL,
    seller_id   TEXT          NOT NULL,
    channel     TEXT          NOT NULL,
    status      TEXT          NOT NULL,
    total_price NUMERIC(12,2) NOT NULL,
    created_at  TIMESTAMPTZ   NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ   NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_orders_buyer_id  ON orders (buyer_id);
CREATE INDEX IF NOT EXISTS idx_orders_seller_id ON orders (seller_id);

CREATE TABLE IF NOT EXISTS order_items (
    id         BIGSERIAL PRIMARY KEY,
    order_id   UUID          NOT NULL REFERENCES orders (id) ON DELETE CASCADE,
    product_id TEXT          NOT NULL,
    name       TEXT          NOT NULL,
    price      NUMERIC(12,2) NOT NULL,
    quantity   INTEGER       NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items (order_id);
