-- migrate:up
CREATE TABLE orders (
    id VARCHAR PRIMARY KEY NOT NULL,
    customer_id VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    status VARCHAR NOT NULL
);

-- migrate:down
DROP TABLE orders;
