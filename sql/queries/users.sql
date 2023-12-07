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
