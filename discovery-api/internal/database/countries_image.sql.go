// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: countries_image.sql

package database

import (
	"context"
)

const createCountryImage = `-- name: CreateCountryImage :one
INSERT INTO countries_image (id, country, country_image, created_at)
VALUES(
gen_random_uuid(), 
$1,
$2,
NOW()
)
On CONFLICT (country)
DO NOTHING
RETURNING id, country, country_image, created_at
`

type CreateCountryImageParams struct {
	Country      string
	CountryImage string
}

func (q *Queries) CreateCountryImage(ctx context.Context, arg CreateCountryImageParams) (CountriesImage, error) {
	row := q.db.QueryRowContext(ctx, createCountryImage, arg.Country, arg.CountryImage)
	var i CountriesImage
	err := row.Scan(
		&i.ID,
		&i.Country,
		&i.CountryImage,
		&i.CreatedAt,
	)
	return i, err
}

const getCountry = `-- name: GetCountry :one
SELECT id, country, country_image, created_at FROM countries_image WHERE country = $1 LIMIT 1
`

func (q *Queries) GetCountry(ctx context.Context, country string) (CountriesImage, error) {
	row := q.db.QueryRowContext(ctx, getCountry, country)
	var i CountriesImage
	err := row.Scan(
		&i.ID,
		&i.Country,
		&i.CountryImage,
		&i.CreatedAt,
	)
	return i, err
}
