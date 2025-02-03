-- name: CreateUserBookmark :one
INSERT INTO users_bookmark (id, user_id, username, place_id, place_name,place_text, created_at)
VALUES(
  gen_random_uuid(),
  $1, $2, $3, $4,$5, NOW()
)
RETURNING *;