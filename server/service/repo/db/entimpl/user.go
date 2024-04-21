package impl

import (
	"context"
	"mobile-telemetry/server/entity"
	"mobile-telemetry/server/service/repo/db/entimpl/ent"

	"github.com/google/uuid"
)

func (db *dbRepo) AddUserIfNotExists(ctx context.Context, user *entity.User) (id uuid.UUID, err error) {
	u, err := db.client.User.Get(ctx, user.ID)
	if err == nil {
		return u.ID, nil
	}

	if !ent.IsNotFound(err) {
		return uuid.UUID{}, err
	}

	u, err = db.client.User.Create().SetID(user.ID).Save(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}

	return u.ID, nil
}
