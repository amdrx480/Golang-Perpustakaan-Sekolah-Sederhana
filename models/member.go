package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name         string         `json:"name"`
	Nis          int            `json:"nis"`
	Gender       string         `json:"gender"`
	Class        int            `json:"class"`
	PlaceOfBirth string         `json:"placeofbirth"`
	DateOfBirth  time.Time      `json:"dateofbirth"`
	PhoneNumber  string         `json:"phonenumber"`
}
