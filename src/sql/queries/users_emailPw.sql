-- name: CreateUserEmailPw :one
INSERT INTO usersEmailPw (id, email, created_at, expired_at, hashed_emailPw, user_id)
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
    hashed_emailPw = EXCLUDED.hashed_emailPw
RETURNING *;

-- name: GetUserEmailPw :one
SELECT * FROM usersEmailPw WHERE hashed_emailPw = $1 AND expired_at > Now() LIMIT 1;
