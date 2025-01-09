// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID          uuid.UUID
	PlaceName   string
	Country     string
	City        string
	Category    string
	PlaceDetail json.RawMessage
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	ID             uuid.UUID
	Username       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Email          string
	HashedPassword string
}
