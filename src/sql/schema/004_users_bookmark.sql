-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users_bookmark (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  username TEXT NOT NULL,
  place_id TEXT NOT NULL,
  place_name TEXT NOT NULL,
  catagory TEXT NOT NULL,
  place_text TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id)REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (place_id)REFERENCES places(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE users_bookmark;