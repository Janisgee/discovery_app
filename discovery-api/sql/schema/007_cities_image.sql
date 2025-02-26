-- +goose Up
CREATE TABLE IF NOT EXISTS cities_image (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  country_id UUID NOT NULL,
  country TEXT NOT NULL,
  city TEXT NOT NULL,
  city_image TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  FOREIGN KEY (country_id)REFERENCES countries_image(id) ON DELETE CASCADE,
  CONSTRAINT unique_country_city UNIQUE (country, city)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_country_city_unique ON cities_image (country, city);

-- +goose Down
DROP TABLE cities_image;