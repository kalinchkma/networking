// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Chrip struct {
	ID        uuid.UUID
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Body      sql.NullString
	UserID    uuid.NullUUID
}

type User struct {
	ID        uuid.UUID
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Email     sql.NullString
}
