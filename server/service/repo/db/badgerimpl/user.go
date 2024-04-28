package impl

import (
	"context"
	"mobile-telemetry/server/entity"
	"mobile-telemetry/server/service/repo/db/badgerimpl/queries"

	"github.com/google/uuid"
)

func (db *dbRepo) AddUserIfNotExists(ctx context.Context, user *entity.User) (id uuid.UUID, err error) {
	tx := db.queries.Update()
	defer tx.Discard()

	err = tx.InsertUser(queries.NewUserKey(user.ID))
	if err != nil {
		return uuid.UUID{}, err
	}

	err = tx.Commit()
	if err != nil {
		return uuid.UUID{}, err
	}

	return user.ID, nil
}
