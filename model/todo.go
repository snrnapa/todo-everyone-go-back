package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	UserId    string    `gorm:"type:varchar(100);not null;primaryKey" json:"user_id"`
	TodoId    string    `gorm:"type:varchar(100);not null;primaryKey" json:"todo_id"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title"`
	Limit     time.Time `gorm:"type:timestamp;" json:"limit"`
	Detail    string    `gorm:"type:varchar(100);" json:"detail"`
	Completed bool      `gorm:"type:boolean;not null" json:"completed"`
	gorm.Model
}
