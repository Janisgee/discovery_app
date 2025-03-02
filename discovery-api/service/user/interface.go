package user

import "github.com/google/uuid"

type UserService interface {
	CreateUser(username string, email string, password string) (*User, error)
	VerifyUserLogin(email string, password string) (*uuid.UUID, error)
	GetUserInfo(email string) (*User, error)
	GetUserProfile(id uuid.UUID) (string, error)
	GetUserPwByID(id *uuid.UUID) (string, error)
	UpdateUserPw(password string, id uuid.UUID) (string, error)
	UpdateUserProfileImage(publicID string, secureURL string, userID *uuid.UUID) (*User, error)
	DisplayUserProfileImage(userID uuid.UUID) (*User, error)
	BeginUserPasswordReset(email string) error
	CompleteUserPasswordReset(resetKey uuid.UUID, newPassword string) error
}
