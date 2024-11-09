-- name: CreateOrder :exec
INSERT INTO orders (id, status, created_at, updated_at, customer_id)
VALUES (@id, @status, @created_at, @updated_at, @customer_id);

-- name: GetOrderForUpdate :one
SELECT *
FROM orders
WHERE id = @order_id
    FOR UPDATE;

-- name: UpdateOrder :exec
UPDATE orders
SET status     = @status,
    created_at = @created_at,
    updated_at = @updated_at
WHERE id = @order_id;