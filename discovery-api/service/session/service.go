package session

import (
	"context"
	"database/sql"
	"discoveryweb/internal/database"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

type postgresSessionService struct {
	db *database.Queries
}

func NewSessionService(db *database.Queries) SessionService {
	return &postgresSessionService{
		db,
	}
}

func (svc *postgresSessionService) CreateSession(userId uuid.UUID) (uuid.UUID, error) {
	// Create an empty context
	ctx := context.Background()

	session, err := svc.db.CreateSession(ctx, userId)
	if err != nil {
		slog.Error("Error when creating new session entry", "err", err)
		return uuid.Nil, err
	}

	return session.ID, nil
}

func (svc *postgresSessionService) CheckAndExtendSession(id uuid.UUID) (uuid.UUID, error) {
	// Create an empty context
	ctx := context.Background()

	session, err := svc.db.ExtendSession(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, NoSessionError
		} else {
			slog.Error("Error when checking/extending session", "err", err)
			return uuid.Nil, err
		}
	}

	return session.UserID, nil
}

func (svc *postgresSessionService) DeleteSession(id uuid.UUID) bool {
	// Create an empty context
	ctx := context.Background()

	err := svc.db.DeleteSession(ctx, id)

	if err != nil {
		slog.Error("Error when deleting session", "err", err)
		return false
	}

	return true
}
