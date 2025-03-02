-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users_password_reset (
  user_id UUID PRIMARY KEY NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  expired_at TIMESTAMP NOT NULL DEFAULT now() + INTERVAL '10 minute',
  reset_key UUID NOT NULL DEFAULT gen_random_uuid(),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX ON users_password_reset(reset_key);

-- +goose Down
DROP TABLE users_password_reset;