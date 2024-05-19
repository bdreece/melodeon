package validator

import "github.com/go-playground/validator/v10"

var Default = &Validator{validator.New()}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i any) error {
	return v.validator.Struct(i)
}
