package models

import "github.com/go-playground/validator/v10"

type BookInput struct {
	Author     string `json:"author" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Publisher  string `json:"publisher" validate:"required"`
	FiscalYear int    `json:"fiscalyear" validate:"required"`
	Isbn       int    `json:"isbn" validate:"required"`
	Qty        int    `json:"qty" validate:"required"`
	Rack       int    `json:"rack" validate:"required"`
}

func (b *BookInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(b)

	return err
}
