package model

import "gorm.io/gorm"

type MstUser struct {
	Id    int    `gorm:"type:numeric(10);unique_index" json:"id"`
	Name  string `gorm:"type:varchar(100);unique_index" json:"name"`
	Age   int    `gorm:"type:numeric(3);unique_index" json:"age"`
	Email string `gorm:"type:varchar(100);unique_index" json:"email"`
	gorm.Model
}
