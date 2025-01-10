package main

import (
	"context"
	"discoveryapp/internal/database"
	"errors"
	"fmt"
)

type UserService interface {
	CreateUser(username string, email string, password string) (*User, error)
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
	_, err := svc.dbQueries.GetUser(ctx, username)
	if err == nil {
		return nil, errors.New("already have the same username saved in database as the input username")
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
