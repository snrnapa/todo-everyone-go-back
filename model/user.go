package model

import "gorm.io/gorm"

type User struct {
	Id       string `gorm:"type:varchar(100);not null;primaryKey" json:"id"`
	Email    string `gorm:"type:varchar(100);not null;uniqueIndex" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Name     string `gorm:"type:varchar(100);" json:"name"`
	Age      int    `gorm:"type:numeric(3)" json:"age"`
	gorm.Model
}
