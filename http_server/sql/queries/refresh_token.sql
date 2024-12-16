-- name: CreateRefreshToken :one
INSERT INTO refresh_token (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetRefreshTokenByToken :one
SELECT * FROM refresh_token WHERE token = $1;

-- name: GetRefreshTokenByUserID :one
SELECT * FROM refresh_token WHERE user_id = $1;