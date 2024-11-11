-- name: GetProductForUpdate :one
SELECT *
FROM products
WHERE id = @product_id
    FOR UPDATE;

-- name: UpdateProduct :exec
UPDATE products
SET name       = @name,
    price      = @price,
    updated_at = @updated_at,
    quantity   = @quantity
WHERE id = @product_id;

-- name: GetProductByID :one
SELECT *
FROM products
WHERE id = @product_id;
