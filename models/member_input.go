package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type MemberInput struct {
	Name         string    `json:"name" validate:"required"`
	Nis          int       `json:"nis" validate:"required"`
	Gender       string    `json:"gender" validate:"required"`
	Class        int       `json:"class" validate:"required"`
	PlaceOfBirth string    `json:"placeofbirth" validate:"required"`
	DateOfBirth  time.Time `json:"dateofbirth" validate:"required"`
	PhoneNumber  string    `json:"phonenumber" validate:"required"`
}

func (b *MemberInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(b)

	return err
}
