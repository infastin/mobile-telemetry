package http

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

var (
	ErrInvalidUUID = validation.NewError("validation_invalid_uuid", "must be a valid UUID")
)

func ValidUUID(value interface{}) error {
	id, ok := value.(uuid.UUID)
	if !ok {
		return errors.New("must be a uuid.UUID")
	}

	if id == (uuid.UUID{}) {
		return ErrInvalidUUID
	}

	return nil
}
