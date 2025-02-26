-- name: CreateCountryImage :one
INSERT INTO countries_image (id, country, country_image, created_at)
VALUES(
gen_random_uuid(), 
$1,
$2,
NOW()
)
On CONFLICT (country)
DO NOTHING
RETURNING *;


-- name: GetCountry :one
SELECT * FROM countries_image WHERE country = $1 LIMIT 1;