package db

import (
	"context"
	"mobile-telemetry/server/entity"

	"github.com/google/uuid"
)

type UserRepo interface {
	AddUserIfNotExists(ctx context.Context, user *entity.User) (id uuid.UUID, err error)
}
