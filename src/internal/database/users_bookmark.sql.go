// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users_bookmark.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUserBookmark = `-- name: CreateUserBookmark :one
INSERT INTO users_bookmark (id, user_id, username, place_id, place_name,place_text, created_at)
VALUES(
  gen_random_uuid(),
  $1, $2, $3, $4,$5, NOW()
)
RETURNING id, user_id, username, place_id, place_name, place_text, created_at
`

type CreateUserBookmarkParams struct {
	UserID    uuid.UUID
	Username  string
	PlaceID   string
	PlaceName string
	PlaceText string
}

func (q *Queries) CreateUserBookmark(ctx context.Context, arg CreateUserBookmarkParams) (UsersBookmark, error) {
	row := q.db.QueryRowContext(ctx, createUserBookmark,
		arg.UserID,
		arg.Username,
		arg.PlaceID,
		arg.PlaceName,
		arg.PlaceText,
	)
	var i UsersBookmark
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Username,
		&i.PlaceID,
		&i.PlaceName,
		&i.PlaceText,
		&i.CreatedAt,
	)
	return i, err
}
