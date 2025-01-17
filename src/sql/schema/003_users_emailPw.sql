-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE usersEmailPw (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email TEXT NOT NULL ,
  created_at TIMESTAMP NOT NULL,
  expired_at TIMESTAMP NOT NULL,
  pw_reset_code TEXT NOT NULL,
  user_id UUID NOT NULL,
  UNIQUE(user_id, email),
  FOREIGN KEY (user_id)REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE usersEmailPw;