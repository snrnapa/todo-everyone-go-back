package model

import "time"

type User struct {
	UserId    string     `gorm:"type:varchar(28);primaryKey" json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
