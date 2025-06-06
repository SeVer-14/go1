-- +goose Up
CREATE TABLE IF NOT EXISTS products
(
    id         SERIAL PRIMARY KEY,
    product_id BIGINT UNIQUE  NOT NULL,
    title      TEXT           NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_products_deleted_at ON products (deleted_at) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS carts
(
    id         SERIAL PRIMARY KEY,
    user_id    BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS cart_items
(
    id         SERIAL PRIMARY KEY,
    cart_id    BIGINT REFERENCES carts (id),
    product_id BIGINT REFERENCES products (product_id),
    quantity   INTEGER NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders
(
    id         SERIAL PRIMARY KEY,
    user_id    BIGINT         NOT NULL,
    status     VARCHAR(20)    NOT NULL DEFAULT 'pending',
    total      DECIMAL(10, 2) NOT NULL,
    cart_id    BIGINT REFERENCES carts (id),
    created_at TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   BIGINT REFERENCES orders (id),
    product_id BIGINT REFERENCES products (product_id),
    quantity   INTEGER        NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для ускорения запросов
CREATE INDEX IF NOT EXISTS idx_carts_user_id ON carts (user_id);
CREATE INDEX IF NOT EXISTS idx_cart_items_cart_id ON cart_items (cart_id);
CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders (user_id);
CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items (order_id);
