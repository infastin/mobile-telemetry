package impl

import (
	"context"
	"mobile-telemetry/server/entity"
	"mobile-telemetry/server/service/repo/db/badgerimpl/schema"

	"github.com/google/uuid"
)

func (db *dbRepo) AddUserIfNotExists(ctx context.Context, user *entity.User) (id uuid.UUID, err error) {
	entry, err := schema.UserEntry(&schema.User{
		ID: user.ID,
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	tx := db.db.NewTransaction(true)
	defer tx.Discard()

	err = tx.SetEntry(entry)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = tx.Commit()
	if err != nil {
		return uuid.UUID{}, err
	}

	return user.ID, nil
}
