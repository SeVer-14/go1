-- +goose Up
CREATE TABLE products
(
    id         SERIAL PRIMARY KEY,
    product_id BIGINT         NOT NULL UNIQUE,
    title      TEXT           NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE carts
(
    id         SERIAL PRIMARY KEY,
    user_id    BIGINT  NOT NULL,
    product_id BIGINT  NOT NULL,
    quantity   INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (product_id) REFERENCES products (product_id)
);

CREATE TABLE orders
(
    id         SERIAL PRIMARY KEY,
    user_id    BIGINT         NOT NULL,
    total      DECIMAL(10, 2) NOT NULL,
    status     VARCHAR(20)    NOT NULL  DEFAULT 'created',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   BIGINT         NOT NULL,
    product_id BIGINT         NOT NULL,
    quantity   INTEGER        NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (product_id) REFERENCES products (product_id)
);
