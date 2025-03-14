// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users_password_reset.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUserPasswordReset = `-- name: CreateUserPasswordReset :one
INSERT INTO users_password_reset (user_id)
VALUES ($1)
ON CONFLICT (user_id)
DO UPDATE
SET
    reset_key = gen_random_uuid(),
    expired_at = now() + INTERVAL '10 minute'
RETURNING user_id, created_at, expired_at, reset_key
`

func (q *Queries) CreateUserPasswordReset(ctx context.Context, userID uuid.UUID) (UsersPasswordReset, error) {
	row := q.db.QueryRowContext(ctx, createUserPasswordReset, userID)
	var i UsersPasswordReset
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.ExpiredAt,
		&i.ResetKey,
	)
	return i, err
}

const deleteUserPasswordReset = `-- name: DeleteUserPasswordReset :exec
DELETE
FROM users_password_reset
WHERE user_id = $1
`

func (q *Queries) DeleteUserPasswordReset(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserPasswordReset, userID)
	return err
}

const getValidUserPasswordReset = `-- name: GetValidUserPasswordReset :one
SELECT user_id, created_at, expired_at, reset_key
FROM users_password_reset
WHERE reset_key = $1
  AND expired_at > Now()
LIMIT 1
`

func (q *Queries) GetValidUserPasswordReset(ctx context.Context, resetKey uuid.UUID) (UsersPasswordReset, error) {
	row := q.db.QueryRowContext(ctx, getValidUserPasswordReset, resetKey)
	var i UsersPasswordReset
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.ExpiredAt,
		&i.ResetKey,
	)
	return i, err
}
