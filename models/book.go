package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Author     string         `json:"author"`
	Title      string         `json:"title"`
	Publisher  string         `json:"publisher"`
	FiscalYear int            `json:"fiscalyear"`
	Isbn       int            `json:"isbn"`
	Qty        int            `json:"qty"`
	Rack       int            `json:"rack"`
}
