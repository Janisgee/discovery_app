package main

import (
	"fmt"
	"log/slog"

	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	// Convert string to []byte
	pass := []byte(password)

	// Hashing the password
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	fmt.Println(string(hash))

	return string(hash), nil
}

func checkPasswordStrength(password string) error {
	// Check password strength
	const minEntropyBits = 60
	err := passwordvalidator.Validate(password, minEntropyBits)
	if err != nil {
		slog.Warn("Password is not strong enough.Plesae signup again with a stronger password. Suggestion:", "error", err)
		return err
	}
	return nil
}
