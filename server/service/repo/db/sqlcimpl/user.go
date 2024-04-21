package impl

import (
	"context"
	"mobile-telemetry/server/entity"

	"github.com/google/uuid"
)

func (db *dbRepo) AddUserIfNotExists(ctx context.Context, user *entity.User) (id uuid.UUID, err error) {
	_, err = db.queries.UpsertUser(ctx, user.ID)
	if err != nil {
		return uuid.UUID{}, err
	}

	return user.ID, nil
}
