-- +goose Up
CREATE TABLE places (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  place_name TEXT NOT NULL UNIQUE,
  country TEXT NOT NULL,
  city TEXT NOT NULL,
  category TEXT NOT NULL,
  place_detail JSONB NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT places_name_country_city_key UNIQUE(place_name, country, city) -- (place_name + country + city) should be unique
);

-- +goose Down
DROP TABLE places;