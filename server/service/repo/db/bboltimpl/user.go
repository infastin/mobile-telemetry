package impl

import (
	"context"
	"mobile-telemetry/server/entity"
	"mobile-telemetry/server/service/repo/db/bboltimpl/queries"

	"github.com/google/uuid"
)

func (db *dbRepo) AddUserIfNotExists(ctx context.Context, user *entity.User) (id uuid.UUID, err error) {
	err = db.queries.InsertUser(queries.NewUserKey(user.ID))
	if err != nil {
		return uuid.UUID{}, err
	}
	return user.ID, nil
}
