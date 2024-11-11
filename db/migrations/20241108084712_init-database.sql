-- migrate:up
CREATE TABLE orders
(
    id          VARCHAR PRIMARY KEY NOT NULL,
    customer_id VARCHAR             NOT NULL,
    created_at  TIMESTAMP           NOT NULL,
    updated_at  TIMESTAMP           NOT NULL,
    status      VARCHAR             NOT NULL
);

CREATE TABLE products
(
    id         VARCHAR PRIMARY KEY NOT NULL,
    name       VARCHAR             NOT NULL,
    price      DECIMAL             NOT NULL,
    created_at TIMESTAMP           NOT NULL,
    updated_at TIMESTAMP           NOT NULL,
    quantity   INT                 NOT NULL
);

CREATE TABLE order_items
(
    id         VARCHAR PRIMARY KEY NOT NULL,
    order_id   VARCHAR             NOT NULL REFERENCES orders (id),
    product_id VARCHAR             NOT NULL REFERENCES products (id),
    quantity   INT                 NOT NULL,
    price      DECIMAL(2)          NOT NULL,
    created_at TIMESTAMP           NOT NULL,
    updated_at TIMESTAMP           NOT NULL
);

-- migrate:down
DROP TABLE order_items;
DROP TABLE orders;
DROP TABLE products;