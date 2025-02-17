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

-- name: GetUserEmailByUsername :one
SELECT email FROM users WHERE id = $1 LIMIT 1;


-- name: GetUserPw :one
SELECT hashed_password FROM users WHERE id = $1 LIMIT 1;

-- name: UpdateUserPw :one
UPDATE users
SET updated_at = NOW(),
hashed_password = $1
WHERE email = $2
RETURNING *;

-- name: UpdateUserPwByID :one
UPDATE users
SET updated_at = NOW(),
hashed_password = $1
WHERE id = $2
RETURNING *;