-- name: CreatePlace :one
INSERT INTO places (id, place_name, country, city, category, place_detail,created_at, updated_at)
VALUES(
gen_random_uuid(),  
$1,
$2,
$3,
$4,
$5,
NOW(),
NOW()
)
RETURNING *;