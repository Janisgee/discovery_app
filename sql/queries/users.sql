-- name: CreateUser :one
ISERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES {
  gen_randowm_uuid(),
  NOW(),
  NOW(),
  $1,
  $2
}
RETURNING *;