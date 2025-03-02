-- name: CreateUserPasswordReset :one
INSERT INTO users_password_reset (user_id)
VALUES ($1)
ON CONFLICT (user_id)
DO UPDATE
SET
    reset_key = gen_random_uuid(),
    expired_at = now() + INTERVAL '10 minute'
RETURNING *;

-- name: GetValidUserPasswordReset :one
SELECT *
FROM users_password_reset
WHERE reset_key = $1
  AND expired_at > Now()
LIMIT 1;

-- name: DeleteUserPasswordReset :exec
DELETE
FROM users_password_reset
WHERE user_id = $1;
