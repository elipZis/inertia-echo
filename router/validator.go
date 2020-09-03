package router

import (
	"github.com/go-playground/validator/v10"
)

// Create a go.validator instance to hook into Echo
type Validator struct {
	validator *validator.Validate
}

//
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

//
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
