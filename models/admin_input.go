package models

import "github.com/go-playground/validator/v10"

type AdminInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,len=6"`
}

func (a *AdminInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(a)

	return err
}
