// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: chirps.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createCirps = `-- name: CreateCirps :one
INSERT INTO chrips (id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING id, created_at, updated_at, body, user_id
`

type CreateCirpsParams struct {
	Body   sql.NullString
	UserID uuid.NullUUID
}

func (q *Queries) CreateCirps(ctx context.Context, arg CreateCirpsParams) (Chrip, error) {
	row := q.db.QueryRowContext(ctx, createCirps, arg.Body, arg.UserID)
	var i Chrip
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.UserID,
	)
	return i, err
}

const getChirps = `-- name: GetChirps :many
SELECT id, created_at, updated_at, body, user_id FROM chrips
`

func (q *Queries) GetChirps(ctx context.Context) ([]Chrip, error) {
	rows, err := q.db.QueryContext(ctx, getChirps)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chrip
	for rows.Next() {
		var i Chrip
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
