package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint      `gorm:"primaryKey"`
	UserId    string    `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title"`
	Deadline  time.Time `gorm:"type:timestamp;" json:"deadline"`
	Detail    string    `gorm:"type:varchar(100);" json:"detail"`
	Completed bool      `gorm:"type:boolean;not null" json:"completed"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Comment struct {
	gorm.Model
	TodoID uint   `gorm:"primaryKey;not null" json:"todo_id"`
	UserId string `gorm:"primaryKey;type:varchar(100);not null" json:"user_id"`
	Text   string `gorm:"type:varchar(255);" json:"text"`
}

type Addition struct {
	TodoID     uint           `gorm:"primaryKey;not null" json:"todo_id"`
	UserId     string         `gorm:"primaryKey;type:varchar(100);not null" json:"user_id"`
	IsFavorite bool           `gorm:"type:boolean;not null" json:"is_favorite"`
	IsBooked   bool           `gorm:"type:boolean;not null" json:"is_booked"`
	IsCheered  bool           `gorm:"type:boolean;not null" json:"is_cheered"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (favo Addition) TableName() string {
	return "trn_additions"
}
