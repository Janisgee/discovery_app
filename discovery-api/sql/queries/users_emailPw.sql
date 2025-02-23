-- name: CreateUserEmailPw :one
INSERT INTO usersEmailPw (id, email, created_at, expired_at, pw_reset_code, user_id)
VALUES(
gen_random_uuid(),
$1,
NOW(),
NOW() + INTERVAL '10 minute',
$2,
$3
)
On CONFLICT (user_id, email)
DO UPDATE
SET 
    created_at = NOW(),
    expired_at = NOW() + INTERVAL '10 minute',
    pw_reset_code = EXCLUDED.pw_reset_code
RETURNING *;

-- name: GetUserEmailPw :one
SELECT * FROM usersEmailPw WHERE pw_reset_code = $1 AND expired_at > Now() LIMIT 1;

-- name: DeleteUserEmailPw :one
DELETE FROM usersEmailPw WHERE pw_reset_code = $1 
RETURNING *;
