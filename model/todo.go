package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	UserId    string    `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title"`
	Limit     time.Time `gorm:"type:timestamp;" json:"limit"`
	Detail    string    `gorm:"type:varchar(100);" json:"detail"`
	Completed bool      `gorm:"type:boolean;not null" json:"completed"`
	gorm.Model
}

type Comment struct {
	gorm.Model
	TodoID uint   `gorm:"not null" json:"todo_id"`
	UserId string `gorm:"type:varchar(100);not null" json:"user_id"`
	Text   string `gorm:"type:varchar(255);" json:"text"`
}
