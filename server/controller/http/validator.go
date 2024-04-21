package http

import "github.com/gookit/validate"

type Validator struct{}

func (Validator) Validate(ptr any) error {
	v := validate.Struct(ptr)
	if !v.Validate() {
		return v.Errors
	}
	return nil
}
