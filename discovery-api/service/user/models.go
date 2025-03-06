package user

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID `json:"id"`
	Username        string    `json:"username"`
	ImagePublicID   string    `json:"public_id"`
	ImageSecureURL  string    `json:"secure_url"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
	Email           string    `json:"email"`
	Hashed_password string    `json:"password"`
}

// UserEmailPw struct
type UserEmailPw struct {
	Email       string    `json:"email"`
	PwResetCode string    `json:"pw_reset_code"`
	UserID      uuid.UUID `json:"user_id"`
}

var ErrNoUser = errors.New("user does not exist")
var ErrPasswordResetInvalid = errors.New("password reset code invalid or expired")
var ErrEmailInUse = errors.New("user email already in use in database")
