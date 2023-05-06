package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type TransactionInput struct {
	BookId    uint      `json:"book_id" validate:"required"`
	MemberID  uint      `json:"member_id" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}

func (b *TransactionInput) Validate() error {
	validate := validator.New()

	//mengembalikan nilai error jika gagal
	err := validate.Struct(b)

	return err
}
