package validator

import "github.com/go-playground/validator/v10"

var v *validator.Validate

func New() {
	v = validator.New()
}

func ValidateStruct(s interface{}) error {
	return v.Struct(s)
}
