// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, username, image_public_id, image_secure_url,created_at, updated_at, email, hashed_password)
VALUES(
gen_random_uuid(),
$1,
$2,
$3,
NOW(),
NOW(),
$4,
$5
)
RETURNING id, username, image_public_id, image_secure_url, created_at, updated_at, email, hashed_password
`

type CreateUserParams struct {
	Username       string
	ImagePublicID  string
	ImageSecureUrl string
	Email          string
	HashedPassword string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.ImagePublicID,
		arg.ImageSecureUrl,
		arg.Email,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ImagePublicID,
		&i.ImageSecureUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, image_public_id, image_secure_url, created_at, updated_at, email, hashed_password FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ImagePublicID,
		&i.ImageSecureUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
	)
	return i, err
}

const getUserEmailByUsername = `-- name: GetUserEmailByUsername :one
SELECT email FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserEmailByUsername(ctx context.Context, id uuid.UUID) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserEmailByUsername, id)
	var email string
	err := row.Scan(&email)
	return email, err
}

const getUserPw = `-- name: GetUserPw :one
SELECT hashed_password FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserPw(ctx context.Context, id uuid.UUID) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserPw, id)
	var hashed_password string
	err := row.Scan(&hashed_password)
	return hashed_password, err
}

const updateUserProfilePicture = `-- name: UpdateUserProfilePicture :one
UPDATE users
SET updated_at = NOW(),
image_public_id = $1,
image_secure_url =$2
WHERE id = $3
RETURNING id, username, image_public_id, image_secure_url, created_at, updated_at, email, hashed_password
`

type UpdateUserProfilePictureParams struct {
	ImagePublicID  string
	ImageSecureUrl string
	ID             uuid.UUID
}

func (q *Queries) UpdateUserProfilePicture(ctx context.Context, arg UpdateUserProfilePictureParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserProfilePicture, arg.ImagePublicID, arg.ImageSecureUrl, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ImagePublicID,
		&i.ImageSecureUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
	)
	return i, err
}

const updateUserPw = `-- name: UpdateUserPw :one
UPDATE users
SET updated_at = NOW(),
hashed_password = $1
WHERE email = $2
RETURNING id, username, image_public_id, image_secure_url, created_at, updated_at, email, hashed_password
`

type UpdateUserPwParams struct {
	HashedPassword string
	Email          string
}

func (q *Queries) UpdateUserPw(ctx context.Context, arg UpdateUserPwParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserPw, arg.HashedPassword, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ImagePublicID,
		&i.ImageSecureUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
	)
	return i, err
}

const updateUserPwByID = `-- name: UpdateUserPwByID :one
UPDATE users
SET updated_at = NOW(),
hashed_password = $1
WHERE id = $2
RETURNING id, username, image_public_id, image_secure_url, created_at, updated_at, email, hashed_password
`

type UpdateUserPwByIDParams struct {
	HashedPassword string
	ID             uuid.UUID
}

func (q *Queries) UpdateUserPwByID(ctx context.Context, arg UpdateUserPwByIDParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserPwByID, arg.HashedPassword, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ImagePublicID,
		&i.ImageSecureUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
	)
	return i, err
}
