package model

import "gorm.io/gorm"

type MstUser struct {
	Id       string `gorm:"type:varchar(100);not null;primaryKey" json:"id"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Name     string `gorm:"type:varchar(100);" json:"name"`
	Age      int    `gorm:"type:numeric(3)" json:"age"`
	Email    string `gorm:"type:varchar(100);not null;primaryKey" json:"email"`
	gorm.Model
}
