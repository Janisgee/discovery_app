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

const deleteUserBookmark = `-- name: DeleteUserBookmark :one
DELETE FROM users_bookmark WHERE place_id = $1 AND user_id = $2 
RETURNING id, user_id, username, place_id, place_name, place_text, created_at
`

type DeleteUserBookmarkParams struct {
	PlaceID string
	UserID  uuid.UUID
}

func (q *Queries) DeleteUserBookmark(ctx context.Context, arg DeleteUserBookmarkParams) (UsersBookmark, error) {
	row := q.db.QueryRowContext(ctx, deleteUserBookmark, arg.PlaceID, arg.UserID)
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

const getAllUserBookmarkPlaceID = `-- name: GetAllUserBookmarkPlaceID :many
SELECT place_id FROM users_bookmark WHERE user_id = $1
`

func (q *Queries) GetAllUserBookmarkPlaceID(ctx context.Context, userID uuid.UUID) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getAllUserBookmarkPlaceID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var place_id string
		if err := rows.Scan(&place_id); err != nil {
			return nil, err
		}
		items = append(items, place_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserBookmark = `-- name: GetUserBookmark :one
SELECT id, user_id, username, place_id, place_name, place_text, created_at FROM users_bookmark WHERE place_id = $1 AND user_id = $2 LIMIT 1
`

type GetUserBookmarkParams struct {
	PlaceID string
	UserID  uuid.UUID
}

func (q *Queries) GetUserBookmark(ctx context.Context, arg GetUserBookmarkParams) (UsersBookmark, error) {
	row := q.db.QueryRowContext(ctx, getUserBookmark, arg.PlaceID, arg.UserID)
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
