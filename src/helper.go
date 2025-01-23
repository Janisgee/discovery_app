package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

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

// Set Cookie
func setSectionCookie(w http.ResponseWriter, token string, expiryTime time.Time) {
	// Set the session id cookie in response, not visible to Javascript (HttpOnly)
	http.SetCookie(w, &http.Cookie{
		Name:     "DA_SESSION_ID",
		Value:    token,
		Expires:  expiryTime,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
	})
}

// Clear the session cookie
func clearSectionCookie(w http.ResponseWriter) {
	// Set the session id cookie in response, not visible to Javascript (HttpOnly)
	http.SetCookie(w, &http.Cookie{
		Name:     "DA_SESSION_ID",
		Value:    "",
		Expires:  time.Unix(0, 0), // Expired immediately
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
}
