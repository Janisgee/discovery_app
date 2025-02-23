package user

import "github.com/google/uuid"

type UserService interface {
	CreateUser(username string, email string, password string) (*User, error)
	VerifyUserLogin(email string, password string) (*uuid.UUID, error)
	GetUserInfo(email string) (*User, error)
	GetUserProfile(id uuid.UUID) (string, error)
	GetUserPwByID(id *uuid.UUID) (string, error)
	CreateUserEmailPw(email string, hashedEmailPw string, user_id uuid.UUID) error
	GetUserEmailFromEmailPw(pwResetCode string) (string, error)
	UpdateUserPw(password string, id uuid.UUID) (string, error)
	ResetUserPw(email string, password string, pw_reset_code string) (*User, error)
	UpdateUserProfileImage(publicID string, secureURL string, userID *uuid.UUID) (*User, error)
	DisplayUserProfileImage(userID uuid.UUID) (*User, error)
}
