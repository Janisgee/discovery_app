-- name: CreateUser :one
INSERT INTO users (id, username, created_at, updated_at, email, hashed_password)
VALUES(
gen_random_uuid(),
$1,
NOW(),
NOW(),
$2,
$3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE email = $1 LIMIT 1;