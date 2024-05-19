package repository

import (
	"github.com/snrnapa/todo-everyone-go-back/db"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Database: db.GetDbInstantce(),
	}
}
func (ur *UserRepository) GetUsers() []model.MstUser {
	var users []model.MstUser
	ur.Database.Find(&users).Order("id desc")
	return users
}
