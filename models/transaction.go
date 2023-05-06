package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Book      Book           `json:"book"`
	BookID    uint           `json:"book_id"`
	Member    Member         `json:"member"`
	MemberID  uint           `json:"member_id"`
}
