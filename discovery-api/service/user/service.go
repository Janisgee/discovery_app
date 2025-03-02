package user

import (
	"context"
	"database/sql"
	"discoveryweb/internal/database"
	"discoveryweb/service/email"
	"discoveryweb/util"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type postgresUserService struct {
	dbQueries *database.Queries
	emailSvc  email.EmailService
}

func NewUserService(dbQueries *database.Queries, emailSvc email.EmailService) UserService {
	return &postgresUserService{
		dbQueries, emailSvc,
	}
}

func (svc *postgresUserService) CreateUser(username string, email string, password string) (*User, error) {
	// Create an empty context
	ctx := context.Background()

	// Check if queries username is in database
	_, err := svc.dbQueries.GetUser(ctx, email)
	if err == nil {
		return nil, errors.New("already have the same email saved in database")
	}

	// Hashed password
	hashedpw, err := util.HashPassword(password)
	if err != nil {
		return nil, errors.New("fail to hash password")
	}

	// Create user if username is not exit in database
	params := database.CreateUserParams{
		Username:       username,
		ImagePublicID:  "https://res.cloudinary.com/dopxvbeju/image/upload/v1740039540/kphottt1vhiuyahnzy8y.jpg",
		ImageSecureUrl: "https://res.cloudinary.com/dopxvbeju/image/upload/v1740039540/kphottt1vhiuyahnzy8y.jpg",
		Email:          email,
		HashedPassword: hashedpw,
	}

	userData, err := svc.dbQueries.CreateUser(ctx, params)
	if err != nil {
		return nil, errors.New("error in creating user into database")
	}

	fmt.Printf("The user created successfully.\n New user data:%v\n", userData)

	// Return the created user
	createdUser := &User{
		Username: userData.Username,
		Email:    userData.Email}

	return createdUser, nil
}

// Get user info by email
func (svc *postgresUserService) GetUserInfo(email string) (*User, error) {
	// Get username from input email
	// Create an empty context
	ctx := context.Background()

	// Check if queries username is in database
	userInfo, err := svc.dbQueries.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoUser
		} else {
			slog.Error("Unexpected error getting user details", "error", err)
			return nil, errors.New("fail to get user info from input email")
		}
	}

	// Return the user info
	userInformation := &User{
		ID:       userInfo.ID,
		Username: userInfo.Username,
		Email:    userInfo.Email}

	return userInformation, nil
}

// Get user email by id
func (svc *postgresUserService) GetUserProfile(id uuid.UUID) (string, error) {
	// Create an empty context
	ctx := context.Background()

	// Get user email from database
	userEmail, err := svc.dbQueries.GetUserEmailByUsername(ctx, id)
	if err != nil {
		slog.Warn("Fail to get user email from database", "error", err)
		return "", errors.New("fail to get user email from database")
	}

	return userEmail, nil
}

// ///////////////////////////////////////////////////////
func (svc *postgresUserService) GetUserPwByID(id *uuid.UUID) (string, error) {
	// Create an empty context
	ctx := context.Background()

	// Get user hashed Pw by ID
	dbhashedPw, err := svc.dbQueries.GetUserPw(ctx, *id)
	if err != nil {
		slog.Warn("error in getting user's email pw information from database", "error", err)
		return "", errors.New(err.Error())
	}
	return dbhashedPw, nil
}

/////////////////////////////////////////////////////////////////////

func (svc *postgresUserService) VerifyUserLogin(email string, password string) (*uuid.UUID, error) {
	// Create an empty context
	ctx := context.Background()

	// Check if queries username is in database
	user, err := svc.dbQueries.GetUser(ctx, email)
	if err != nil {
		return nil, errors.New("cannot find input email from database")
	}

	// Valified database hashed password and user input hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return nil, errors.New("fail to login as password is not correct. Please try again")
	}

	return &user.ID, nil
}

func (svc *postgresUserService) UpdateUserPw(password string, id uuid.UUID) (string, error) {
	// Create an empty context
	ctx := context.Background()

	// Hashed password
	hashedpw, err := util.HashPassword(password)
	if err != nil {
		return "Fail to update user password.", errors.New("fail to hash password")
	}

	// Params to update user
	params := database.UpdateUserPwByIDParams{
		HashedPassword: hashedpw,
		ID:             id,
	}

	// Update User Pw
	userInfo, err := svc.dbQueries.UpdateUserPwByID(ctx, params)
	if err != nil {
		return "Fail to update user password.", errors.New("already have the same email saved in database")
	}
	fmt.Printf("The user password is successfully updated.\n User password updated at:%v\n", userInfo.UpdatedAt)

	return "Successfully updated user password.", nil
}

func (svc *postgresUserService) UpdateUserProfileImage(publicID string, secureURL string, userID *uuid.UUID) (*User, error) {
	// Create an empty context
	ctx := context.Background()

	// Update user password in user database
	params := database.UpdateUserProfilePictureParams{
		ImagePublicID:  publicID,
		ImageSecureUrl: secureURL,
		ID:             *userID,
	}
	userInfo, err := svc.dbQueries.UpdateUserProfilePicture(ctx, params)
	if err != nil {
		return nil, errors.New("error in updating user new profile picture")
	}

	fmt.Printf("The user profile picture is successfully updated.\n User password updated at:%v\n", userInfo.UpdatedAt)

	// Return the user info
	userInformation := &User{
		ID:             userInfo.ID,
		ImagePublicID:  userInfo.ImagePublicID,
		ImageSecureURL: userInfo.ImageSecureUrl}

	return userInformation, nil
}

func (svc *postgresUserService) DisplayUserProfileImage(userID uuid.UUID) (*User, error) {
	// Create an empty context
	ctx := context.Background()

	// Get user image information
	userInfo, err := svc.dbQueries.GetUserProfileImageInfo(ctx, userID)
	if err != nil {
		return nil, errors.New("error in updating user new profile picture")
	}

	// Return the user info
	userInformation := &User{

		ImagePublicID:  userInfo.ImagePublicID,
		ImageSecureURL: userInfo.ImageSecureUrl}

	return userInformation, nil
}

func (svc *postgresUserService) BeginUserPasswordReset(email string) error {
	// Create an empty context
	ctx := context.Background()

	user, err := svc.GetUserInfo(email)

	if err != nil {
		return err
	}

	reset, err := svc.dbQueries.CreateUserPasswordReset(ctx, user.ID)

	if err != nil {
		slog.Error("failed to insert user password reset code", "err", err)
		return err
	}

	// TODO: ENV var for site base url.
	resetLink := "http://mysite.com/forgetPwChange?" + "evpw=" + reset.ResetKey.String()

	err = svc.emailSvc.SendPasswordResetEmail(user.Username, email, resetLink)
	if err != nil {
		slog.Error("failed to send password reset email", "err", err)
		return err
	}

	return nil
}

func (svc *postgresUserService) CompleteUserPasswordReset(resetKey uuid.UUID, newPassword string) error {
	// Create an empty context
	ctx := context.Background()

	resetInfo, err := svc.dbQueries.GetValidUserPasswordReset(ctx, resetKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrPasswordResetInvalid
		} else {
			slog.Error("Unexpected error resetting password", "err", err)
			return err
		}
	}

	hashedPassword, err := util.HashPassword(newPassword)
	if err != nil {
		slog.Warn("failed to hash password", "err", err)
		return err
	}

	_, err = svc.dbQueries.UpdateUserPwByID(ctx, database.UpdateUserPwByIDParams{ID: resetInfo.UserID, HashedPassword: hashedPassword})
	if err != nil {
		slog.Error("failed to update user password on reset", "err", err)
		return err
	}

	err = svc.dbQueries.DeleteUserPasswordReset(ctx, resetInfo.UserID)
	if err != nil {
		slog.Error("failed to delete user password reset after reset complete", "err", err)
		// No need to return error here - we've already completed the password reset so just ignore this.
	}
	return nil
}
