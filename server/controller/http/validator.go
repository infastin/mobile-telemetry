package http

import (
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

type Validator struct{}

func (Validator) Validate(ptr any) error {
	v := validate.Struct(ptr)
	if !v.Validate() {
		return v.Errors
	}
	return nil
}

func init() {
	validate.AddValidator("required_uuid", func(id uuid.UUID) bool {
		return id != uuid.UUID{}
	})
}
