-- name: CreateUser :one
INSERT INTO users
(id, created_at, updated_at, deleted_at, email, phone, first_name, last_name, username, password)
VALUES
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdatePassword :one
UPDATE users
SET password = $2
WHERE id = $1
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 AND deleted_at IS NULL;

-- name: DeleteUserByID :one
UPDATE users
SET deleted_at = $2
WHERE id = $1
RETURNING *;
