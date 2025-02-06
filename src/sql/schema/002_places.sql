-- +goose Up
CREATE TABLE places (
  id TEXT PRIMARY KEY UNIQUE,
  place_name TEXT NOT NULL UNIQUE,
  country TEXT NOT NULL,
  city TEXT NOT NULL,
  category TEXT NOT NULL,
  place_detail JSONB NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT places_id_place_name_country_city_key UNIQUE(id, place_name, country, city) -- (place_name + country + city) should be unique
);

-- +goose Down
DROP TABLE places;