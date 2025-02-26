-- +goose Up
CREATE TABLE countries_image (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  country TEXT NOT NULL UNIQUE,
  country_image TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now()
);


-- +goose Down
DROP TABLE countries_image;