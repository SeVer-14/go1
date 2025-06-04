-- +goose Up
CREATE TABLE products
(
    id         SERIAL PRIMARY KEY,
    product_id BIGINT         NOT NULL UNIQUE,
    title      TEXT           NOT NULL,
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

