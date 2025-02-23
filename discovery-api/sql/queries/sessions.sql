-- name: CreateSession :one
INSERT INTO sessions (user_id)
VALUES($1)
RETURNING *;

-- name: GetSession :one
SELECT id, user_id, created_at, expires_at
FROM sessions
WHERE id = $1
ORDER BY expires_at DESC
LIMIT 1;

-- name: ExtendSession :one
UPDATE sessions
SET expires_at = now() + INTERVAL '30 minutes'
WHERE id = $1 AND expires_at > now()
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1;