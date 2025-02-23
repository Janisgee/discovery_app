-- name: CreatePlace :one
INSERT INTO places (id, place_name, country, city, category, place_detail,created_at, updated_at)
VALUES(
$1,  
$2,
$3,
$4,
$5,
$6,
NOW(),
NOW()
)
On CONFLICT (id, place_name,country, city)
DO NOTHING
RETURNING *;

-- name: GetPlace :one
SELECT * FROM places WHERE id = $1 LIMIT 1;


-- name: GetPlaceIDFromDB :one
SELECT id FROM places WHERE place_name = $1 LIMIT 1;