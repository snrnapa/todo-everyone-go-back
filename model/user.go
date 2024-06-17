package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId    string     `gorm:"type:varchar(28);primaryKey" json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type Contact struct {
	ID        uint           `gorm:"primaryKey"`
	UserId    string         `gorm:"type:varchar(28)" json:"user_id"`
	Text      string         `gorm:"type:varchar(200)" json:"text"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
