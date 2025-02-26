-- name: CreateCityImage :one
INSERT INTO cities_image (id, country_id, country, city, city_image, created_at)
VALUES(
gen_random_uuid(), 
$1,
$2,
$3,
$4,
NOW()
)
On CONFLICT (country, city)
DO NOTHING
RETURNING *;

-- name: GetCity :one
SELECT * FROM cities_image WHERE city=$1 AND country=$2 LIMIT 1;