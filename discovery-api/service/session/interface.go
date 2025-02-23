package session

import (
	"github.com/google/uuid"
)

type SessionService interface {
	CreateSession(userId uuid.UUID) (uuid.UUID, error)
	CheckAndExtendSession(id uuid.UUID) (uuid.UUID, error)
	DeleteSession(id uuid.UUID) bool
}
