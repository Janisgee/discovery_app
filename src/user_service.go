package main

import (
	"context"
	"discoveryapp/internal/database"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(username string, email string, password string) (*User, error)
	VerifyUserLogin(email string, password string) (*uuid.UUID, error)
	GetUserInfo(email string) (*User, error)
}

// User struct to hold input data
type User struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Hashed_password string `json:"password"`
}

type PostgresUserService struct {
	dbQueries *database.Queries
}

func (svc *PostgresUserService) CreateUser(username string, email string, password string) (*User, error) {
	// Create an empty context
	ctx := context.Background()

	// Check if queries username is in database
	_, err := svc.dbQueries.GetUser(ctx, email)
	if err == nil {
		return nil, errors.New("already have the same email saved in database")
	}

	// Hashed password
	hashedpw, err := hashPassword(password)
	if err != nil {
		return nil, errors.New("fail to hash password")
	}

	// Create user if username is not exit in database
	params := database.CreateUserParams{
		Username:       username,
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

// Get user info
func (svc *PostgresUserService) GetUserInfo(email string) (*User, error) {
	// Get username from input email
	// Create an empty context
	ctx := context.Background()

	// Check if queries username is in database
	userInfo, err := svc.dbQueries.GetUser(ctx, email)
	if err != nil {
		slog.Warn("Fail to get user info from input email", "error", err)
		return nil, errors.New("fail to get user info from input email")
	}

	fmt.Printf("The retrieved user data:%v\n", userInfo)

	// Return the user info
	userInformation := &User{
		Username: userInfo.Username,
		Email:    userInfo.Email}

	return userInformation, nil
}

/////////////////////////////////////////////////////////////////////

func (svc *PostgresUserService) VerifyUserLogin(email string, password string) (*uuid.UUID, error) {
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
