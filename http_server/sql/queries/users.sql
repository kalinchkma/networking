-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;


-- name: GetUser :many
SELECT * FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=$1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id=$1;

-- name: UpdateUserByID :exec
UPDATE users SET email = $1, hashed_password = $2 WHERE id = $3;